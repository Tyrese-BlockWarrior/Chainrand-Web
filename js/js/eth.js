var ERROR_NO_ETH = 101;
var ERROR_UNSUPPORTED_CHAIN = 102;
var ERROR_EXECUTION = 103;
var ERROR_LOGIN_PENDING = 105;

var CONTRACTS_CONFIG_URL = 'contracts/config.json';

var W = {
    walletAddress: null, 
    config: null,
    contracts: null
};

var setupEth = function (W, f) {

    W.contracts = {};
    for (var k in W.config) {
        W.contracts[k] = {};
        for (var c in W.config[k])
            W.contracts[k][c] = null;
    }

    var localstoragePrefix = 'crls_';    
    var localstorageGet = function (key, defaultValue) {
        var k = localstoragePrefix + key;
        var c = window.localStorage.getItem(k);
        if (c === null) {
            window.localStorage.setItem(k, defaultValue);
            return defaultValue;
        } 
        return c;
    };

    var localstorageSet = function (key, value) {
        var k = localstoragePrefix + key;
        window.localStorage.setItem(k, value);
    };

    var tryExec = function (f, k) {
        var fn = f ? f[k] : null, args = [];
        for (var i = 2; i < arguments.length; ++i) 
            args.push(arguments[i]);
        if ($.isFunction(fn)) {
            fn.apply(undefined, args);

        } else if (k == 'success' && $.isFunction(f)) {
            f.apply(undefined, args);
        }
    };
    
    var getEthErrorMessage = function (e) {
        W.ethError = e;
        var q = '' + (e ? e.message : '');
        if ((/insufficient/i).test(q)) 
            return "Insufficient funds for payment and gas fees.";
        try {
            var s = JSON.parse(q.substr(q.indexOf('{'))), v;
            if ( ((v = s.originalError) && (v = v.message)) ||
                 (v = s.message) ) {
                s = v.replace(/^[^:]+:\s*/, '');    
            } 
            return s;
        } catch (error) {
            return 'Transaction failed.';
        }
    };

    if (typeof window.ethereum !== 'undefined') {
        W.web3 = new Web3(window.ethereum);
    }

    var chainIsSupported = function (chainId) {
        for (var k in W.config) {
            if (!W.config[k].addresses[chainId])
                return false;
        }
        return true;
    }

    var extend = function(f, success) {
        return { success: success, error: function(e) { tryExec(f, 'error', e) } }
    }

    var callChainId = function (f) {
        if (W.web3) {
            W.web3.eth.net.getId().then(function(chainId) {
                if (chainIsSupported(chainId)) {
                    W.chainId = chainId;
                    W.chainUnsupported = 0;
                    tryExec(f, 'success', chainId);
                } else {
                    W.chainId = null;
                    W.walletAddress = null;
                    W.chainUnsupported = 1;
                    tryExec(f, 'error', {code: ERROR_UNSUPPORTED_CHAIN});
                }
            }).catch(function (error) {
                console.log(error);
                W.chainId = null;
                W.walletAddress = null;
                W.chainUnsupported = 0;
                tryExec(f, 'error', {code: ERROR_NO_ETH});
            });
        } else {
            W.chainId = null;
            W.walletAddress = null;
            W.chainUnsupported = 0;
            tryExec(f, 'error', {code: ERROR_NO_ETH});
        }
    }

    var signIn = function (f, silent) {
        console.log('Signing in wallet ' + (silent ? '(silent)' : '') + '...');
        callChainId(extend(f, function (chainId) {
            console.log('Signed in: ' + chainId);
            W.chainId = chainId;
            var p = silent ? W.web3.eth.getAccounts() :
                window.ethereum.request({method: 'eth_requestAccounts'});
            
            p.then(function (accounts) {
                W.walletAddress = accounts[0];
                tryExec(f, 'success', W.walletAddress);    
            }).catch(function (error) {
                W.walletAddress = null;
                // Handle error. Likely the user rejected the login
                tryExec(f, 'error', {code: ERROR_LOGIN_PENDING});
            })
        }))
    }

    var silentSignIn = function (f) {
        signIn(f, 1);
    }

    var isSignedIn = function () {
        return W.walletAddress !== null && W.chainId !== null
    }

    var requireSignIn = function (f) {
        if (isSignedIn()) {
            tryExec(f, 'success')
        }
    }

    var callContract = function (name) {
        var a = arguments;
        var n = a.length;
        var f = a[n - 1];
        var methodName = n > 2 ? a[1] : '';
        var methodArgs = [];
        if (n > 3) {
            for (var i = 2; i < n - 1; ++i) {
                methodArgs.push(a[i]);
            }
        }
        var nameComps = name.split('!');
        name = nameComps[0];

        var ff = extend(f, function (chainId) {
            if (!W.contracts[name][chainId]) {
                W.contracts[name][chainId] = new W.web3.eth.Contract(
                    W.config[name].abi, 
                    W.config[name].addresses[chainId]);
            }
            if (methodName == '') {
                tryExec(f, 'success', W.contracts[name][chainId]);    
            } else {
                try {
                    W.contracts[name][chainId].methods[methodName]
                    .apply(undefined, methodArgs).call()
                    .then(function (x) {
                        tryExec(f, 'success', x);
                    }).catch(function (error) {
                        tryExec(f, 'error', {code: ERROR_EXECUTION});
                    })     
                } catch (e) {
                    tryExec(f, 'error', {code: ERROR_EXECUTION});
                }
            }
        });

        if (nameComps.length > 1) {
            silentSignIn(extend(ff, function () {
                if (W.isSignedIn()) {
                    tryExec(ff, 'success', W.chainId);
                } else {
                    signIn(extend(ff, function () {
                        if (W.walletAddress && W.chainId) {
                            tryExec(ff, 'success', W.chainId);
                        } else {
                            tryExec(ff, 'error', {code: ERROR_NO_ETH});
                        }
                    }))
                }
            }))
        } else {
            callChainId(ff)    
        }

    }

    var authorizeLinkPendingPrefix = 'authorizeLinkPending';

    var authorizeLinkPendingAdd = function (txHash) {
        var p = localstorageGet(authorizeLinkPendingPrefix, {});
        p[txHash] = {
            walletAddress: W.walletAddress, 
            chainId: W.chainId
        };
        localstorageSet(authorizeLinkPendingPrefix, p);
    }
    
    var authorizeLinkPendingFirst = function () {
        var p = localstorageGet(authorizeLinkPendingPrefix, {});
        for (var k in p) {
            var h = p[k];
            if (h.walletAddress == W.walletAddress && 
                h.chainId == W.chainId)
                return k;
        }
        return null;
    }

    var authorizeLinkPendingDelete = function (txHash) {
        var p = localstorageGet(authorizeLinkPendingPrefix, {});
        delete p[txHash];
        localstorageSet(authorizeLinkPendingPrefix, p);   
    }

    var waitTransactionMined = function (txHash, f) {
        W.web3.eth.getTransactionReceipt(txHash, function (error, receipt) {
            if (error) {
                // do nothing
            } else if (receipt == null) {
                setTimeout(function () {
                    waitTransactionMined(txHash, f);
                }, 5000);
            } else {
                tryExec(f, 'success', receipt.status);
            }
        });
    }

    var checkHasAuthorizedLink = function (f) {
        callContract('chainlink!', 
            'allowance', 
            W.walletAddress, 
            W.config['chainrand'].addresses[W.chainId], 
            extend(f, function (x) {
                x = '' + x;
                if (x != '0') {
                    tryExec(f, 'success', true);
                } else {
                    var txHash = authorizeLinkPendingFirst();
                    if (txHash) {
                        tryExec(f, 'pending');    
                        waitTransactionMined(txHash, function (status) {
                            authorizeLinkPendingDelete(txHash);
                            tryExec(f, 'success', status);    
                        })
                    } else {
                        tryExec(f, 'success', false);    
                    }
                }
            })
        )
    }
    
    var estimateGasAndSend = function (m, p, f) {
        m.estimateGas(p)
        .then(function (e) {
            p.gas = e;
            m.send(p) 
            .on('transactionHash', function (txHash) {
                tryExec(f, 'success', txHash);
            }).catch(function (e) {
                console.log(e);
                tryExec(f, 'error', {code: ERROR_EXECUTION, message: 'Transaction failed.'});
            });
        }).catch(function (e) {
            tryExec(f, 'error', {code: ERROR_EXECUTION, message: getEthErrorMessage(e)});
        })
    }

    var authorizeLink = function (f) {
        callContract('chainlink!', extend(f, function (cl) {
            var chainrandAddress = W.config['chainrand'].addresses[W.chainId];
            var amountToApprove = '100000000000000000000000000';
            var m = cl.methods['approve'](chainrandAddress, amountToApprove);
            estimateGasAndSend(m, {from: W.walletAddress}, extend(f, function (txHash) {
                tryExec(f, 'pending', txHash);
                authorizeLinkPendingAdd(txHash);
                waitTransactionMined(txHash, function (status) {
                    authorizeLinkPendingDelete(txHash);
                    tryExec(f, 'success', status);    
                })
            }));
        }))
    }

    // function mint(string memory _name, uint _seedKeyHash, uint _codeHash, 
    // string memory _codeURI, string memory _imageURI, string memory _projectURI) 

    var mint = function (name, seedKeyHash, codeHash, codeURI, imageURI, projectURI, value, f) {
        callContract('chainrand!', extend(f, function (cr) {
            var m = cr.methods['mint'](name, seedKeyHash, codeHash, codeURI, imageURI, projectURI);
            estimateGasAndSend(m, {from: W.walletAddress, value: value}, f); 
        }))
    }

    var updateToken = function (tokenId, key, value, f) {
        var methodName = 'set' + capitalize(key);
        callContract('chainrand!', extend(f, function (cr) {
            var m = cr.methods[methodName](tokenId, value);
            estimateGasAndSend(m, {from: W.walletAddress}, f); 
        }))
    }

    var payToken = function (tokenId, value, f) {
        callContract('chainrand!', extend(f, function (cr) {
            var m = cr.methods['pay'](tokenId);
            estimateGasAndSend(m, {from: W.walletAddress, value: value}, f); 
        }))   
    }

    var getEtherscanLink = function (contractName) {
        var p = '';
        if (W.chainId == 80001) {
            p = 'https://mumbai.polygonscan.com/address/';
        }
        if (W.chainId == 1) {
            p = 'https://etherscan.io/address/';        
        }
        if (W.chainId == 4) {
            p = 'https://rinkeby.etherscan.io/address/';
        }
        if (W.chainId == 137) {
            p = 'https://polygonscan.com/address/';
        }   
        if (W.chainId) {
            var v; 
            if ((v = W.config[contractName]) && (v = v.addresses) && (v = v[W.chainId]))
                return p + v;
        }
        return '';
    }
    W.getEtherscanLink = getEtherscanLink;
    W.signIn = signIn;
    W.silentSignIn = silentSignIn;
    W.chainIsSupported = chainIsSupported;
    W.callContract = callContract;
    W.authorizeLink = authorizeLink;
    W.checkHasAuthorizedLink = checkHasAuthorizedLink;
    W.waitTransactionMined = waitTransactionMined;
    W.mint = mint;
    W.updateToken = updateToken;
    W.payToken = payToken;
    W.isSignedIn = isSignedIn;

    window.ethereum.on('accountsChanged', function (accounts) {
        silentSignIn(f);
    });
    window.ethereum.on('networkChanged', function (networkId) {
        silentSignIn(f);
    });
    window.ethereum.on('chainChanged', function (networkId) {
        silentSignIn(f);
    });
    
    silentSignIn(f);
};