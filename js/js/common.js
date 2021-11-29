var chainlist = [
	{name: 'mainnet', currency: 'ETH', id: 1, supported: false},
	{name: 'rinkeby', currency: 'ETH', id: 4, supported: true},
	{name: 'polygon', currency: 'MATIC', id: 137, supported: false},
	{name: 'mumbai', currency: 'MATIC', id: 80001, supported: true},
];

var githubURL = 'https://github.com/chainrand';

var tokensSyncCounter = 0;


var h2d = function (s) {

	var add = function (x, y) {
		var c = 0, r = [];
		var x = x.slice();
		var y = y.slice();
		while (x.length || y.length) {
			var s = (x.pop() || 0) + (y.pop() || 0) + c;
			r.unshift(s < 10 ? s : s - 10); 
			c = s < 10 ? 0 : 1;
		}
		if (c) r.unshift(c);
		return r;
	}

	s = s.trim()
	if ((/^(0x)?[0-9a-f]+$/i).test(s)) {
		var dec = [0];
		for (var i = 0; i < s.length; ++i) {
			var n = parseInt(s.charAt(i), 16);
			for (var t = 8; t; t >>= 1) {
				dec = add(dec, dec);
				if (n & t) dec = add(dec, [1]);
			}
		}
		return dec.join('');	
	}
	return null    
}

var getVW = function () { return Math.max(document.documentElement.clientWidth, window.innerWidth || 0) };
var getVH = function () { return Math.max(document.documentElement.clientHeight, window.innerHeight || 0) };

var n = navigator.userAgent;
var isNativeAndroid = n.match(/Linux/i) && n.match(/Android/i) && n.match(/Mozilla/i) && n.match(/AppleWebKit/i) && n.match(/Chrome/i);
var isIOS = n.match(/iPhone|iPad|iPod/i) && n.match(/AppleWebKit/);
var isMobile = n.match(/iPad|iPhone|Android|BlackBerry|webOS|Mobile/i);
var isPWA = isMobile && ((window.matchMedia('(display-mode: standalone)').matches) || (window.navigator.standalone));

var capitalize = function (s) { s = '' + s; return s.charAt(0).toUpperCase() + s.slice(1); };

var throttle = function (fn, wait) {
	var scheduled = 0;
	var ex = function () { scheduled = 0; fn(); };
	return function() {
		if (!scheduled) {
			scheduled = 1;
			setTimeout(ex, wait);
		}
	}
}

var lowercaseEquals = function (a, b) {
	return a.toLowerCase().trim() == b.toLowerCase().trim();
}

var base36Counter = function (x) {
	var state = (x || '1').split('');
	var chars = '0123456789abcdefghijklmnopqrstuvwxyz';
	return function (increment) {
		var d = increment > 0 ? 1 : -1;
		for (var j = Math.abs(increment); j--;) 
			for (var i = state.length; i--;) {
				var p = chars.indexOf(state[i]) + d;
				if (p < chars.length) {
					state[i] = chars.charAt(p);
					i = 0;
				} else {
					state[i] = '0';
					if (!i) {
						state.unshift('1');
						i = 0;
					}
				}
			}	
		return state.join('');
	};
};

var elId = (function () {
	var b = 'elId_', i = base36Counter();
	return function () {
		return b + i(1);
	};
})();

var walletConnector = (function () {
	var self = {};
	var ids = {};

	var getId = function (key) {
		if (ids[key]) return ids[key];
		ids[key] = elId();
		return ids[key];
	}

	self.ctx = function (key) {
		var id = getId(key);
		return ['div', {class: 'wallet', id: id}];
	}

	self.sync = function (key) {
		var id = getId(key);
		var a = W.walletAddress;
		var ctx = ['div', {class: 'info'}];
		if (a) {
			a = '0x' + (a.substr(2, 6) + '...' + a.substr(-6)).toUpperCase();
			a = ['span', ['i', {class: 'icon-link'}], ' ', a];
		} else if (W.chainUnsupported) {
			a = ['span', ['i', {class: 'icon-attention'}]];
		} else {
			a = ['span', ['i', {class: 'icon-login'}], ' Connect Wallet '];
			ctx.push({class: 'unconnected'});
		}
		ctx.push(['span', {class: 'address'}, a]);

		if (W.chainId) {
			var chainName = '';
			for (var i = 0; i < chainlist.length; ++i) {
				if (chainlist[i].id == W.chainId)
					chainName = chainlist[i].name;
			}
			ctx.push(['span', {class: 'c-chain'},
				' (', ['span', {class: 'chain'}, capitalize(chainName)], ')'
			]);
		} else if (W.chainUnsupported) {
			ctx.push(['span', {class: 'c-chain'},
				' Chain Unsupported!'
			]);
		}


		$('#' + id).html(HTML(ctx)).find('.info').click(function () {
			if (!W.walletAddress && !W.chainUnsupported) {
				W.signIn(function () {
					Layout.sync()	
				});
			} 
		});
	}
	return self;
})()

var validateCID = function (s) {
	return /^(Qm[1-9A-HJ-NP-Za-km-z]{44,}|b[A-Za-z2-7]{58,}|B[A-Z2-7]{58,}|z[1-9A-HJ-NP-Za-km-z]{48,}|F[0-9A-F]{50,})$/.test(s);
}

var validateURL = function (s) {
	return /^(?:(?:(?:https?|ftp):)?\/\/)(?:\S+(?::\S*)?@)?(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)(?:\.(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)*(?:\.(?:[a-z\u00a1-\uffff]{2,})))(?::\d{2,5})?(?:[/?#]\S*)?$/i.test(s);
}

var validateSHA256Hex = function (s) {
	return /^(0x)?[0-9a-f]{64}$/i.test(s);
}

var validateDecimalPositive = function (s) {
	return /^(0?\.0*?[0-9]*|[1-9][0-9]*\.?[0-9]+|0)$/.test(s)
}

var gweiPerEth = new Big('1000000000000000000');

var decimalToGwei = function (decimal) {
	try {
		var p = new Big('' + decimal);
		return p.mul(gweiPerEth) + '';
	} catch (e) {
		return '';
	}
}

var gweiToDecimal = function (gwei) {
	try {
		var p = new Big('' + gwei);
		return p.div(gweiPerEth) + '';
	} catch (e) {
		return '';
	}	
}

var sha256Hex = function (s) {
	var hash = sha256.create();
	hash.update(s);
	return hash.hex();
}

var enableElement = function ($el) {
	$el.removeAttr('disabled');
}

var disableElement = function ($el) {
	$el.attr('disabled', 'disabled');
}

var genOpenSeaURL = function (tokenId, chainId) {
	var p = '';
	var mainnetPath = 'https://opensea.io/assets';
	var testnetPath = 'https://testnets.opensea.io/assets';
	if (a == 1) {
		p = mainnetPath;
	} else if (chainId == 4) {
		p = testnetPath;
	} else if (chainId == 137) {
		p = mainnetPath + '/matic';
	} else if (chainId == 80001) {
		p = testnetPath + '/mumbai';
	}
	if (p) {
		var a; 
		if ((a = W.config) && (a = a['chainrand']) && 
			(a = a.addresses) && (a = a[chainId])) {
			return p + '/' + a + '/' + tokenId; 
		}	
	}
	return '';
}

var trim = function (s) { 
	return String.prototype.trim ? s.trim() : s.replace(/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g, ''); 
};

var isNonEmptyString = function (s) {
	if (typeof s === 'string' || f.constructor == String)
		return s.length > 0;
}

$(function () {
	var ma = modal.alert;
	var newAlert = function (msg) {
		var bodyCtx = ['div', {class: 'alert-body'}, msg];
		var footerCtx = ['div', {class: 'alert-footer'}, 
			['button', {class: 'ok light'}, 'Ok']
		];
		ma({
			body: HTML(bodyCtx), 
			footer: HTML(footerCtx), 
			afterShow: function () {
				$('.alert-footer .ok').click(function () {
					ma.close();
				})
			}
		})
	};
	var newConfirm = function (msg, onConfirm) {
		var bodyCtx = ['div', {class: 'alert-body'}, msg];
		var footerCtx = ['div', {class: 'alert-footer'}, 
			['button', {class: 'cancel light'}, 'Cancel'],
			['div', {class: 'spacer'}],
			['button', {class: 'ok'}, 'Ok']
		];
		ma({
			body: HTML(bodyCtx), 
			footer: HTML(footerCtx), 
			afterShow: function () {
				$('.alert-footer .ok').click(function () {
					ma.close();
					onConfirm();
				})
				$('.alert-footer .cancel').click(function () {
					ma.close();
				})
			}
		})
	}
	modal.alert = newAlert;
	modal.confirm = newConfirm;
})

var toObject = function (data) {
	var o = {};
	for (var k in data) {
		var v = data[k].value;
		if (isNonEmptyString(v))
			v = trim(v);
		o[data[k].name] = v;
	}
	return o;
}

var setupErrorHideOnChange = function ($input, $error) {
	var events = "input,keydown,keyup,mousedown,mouseup,select,contextmenu,drop,change".split(',');
	for (var i = 0; i < events.length; ++i) 
		$input.on('change', function () {
			var currValue = $input.val();
			var prevValue = $input.attr('_prev');
			if (prevValue != currValue) {
				if ($error.is(':visible')) {
					$error.hide();
				}
			}
			$input.attr('_prev', currValue);
		});
}

var scrollToTop = function () {
	$('html, body').animate({
        scrollTop: 0
    }, 500);
}