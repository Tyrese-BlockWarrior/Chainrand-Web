(function () {
	
	var genFieldCtx = function (label, name, required, description) {
		return ['div', {class: 'f-row', name: name},
			['div', {class: 'label'}, 
				label, (required? ' *': ''), 
				['div', {class: 'help', name: name}, ['div', '?']]
			],
			['div', {class: 'field'},
				['div', {class: 'error'}, 'Lorem Ipsum'],
				['input', {name: name, type: 'text'}],
			],
			(description ? ['div', {class: 'description', name: name}, description] : '')
		];
	}

	var paymentLabel = ['span', 
		'Payment (', 
			['span', {class: 'currency'}, 'ETH'], 
			['span', {class: 'min'}, ', min: ', ['span', {class: 'amount'}, '0.0']], 
		')'
	];

	var authorizeLinkText = 'Authorize LINK payment';

	var $page, $authorize, $mint, $form;

	var setupOnChange = function (name) {
		var $r = $form.find('.f-row[name="' + name + '"]');
		setupErrorHideOnChange($r.find('input'), $r.find('.error'));
	}

	var validate = function (data) {
		
		var errors = [];

		var assertNotEmpty = function (name) {
			if (data[name].length < 1) 
				errors.push({name: name, message: 'Cannot be empty!'});
		}
		var assertValidURI = function (name) {
			var value = data[name];
			if (!validateURL(value) && !validateCID(value)) 
				errors.push({name: name, message: 'Must be a valid URL or CID!'});
		}

		assertNotEmpty('name');
		assertNotEmpty('seedKey');
		assertNotEmpty('codeHash');
		
		if (!validateSHA256Hex(data['codeHash'])) 
			errors.push({name: 'codeHash', message: 'Not a valid SHA-256 hash!'});

		assertNotEmpty('codeURI');
		assertValidURI('codeURI');

		if (data['imageURI']) 
			assertValidURI('imageURI');

		if (data['projectURI']) 
			assertValidURI('projectURI');

		if (data['payment'] && !validateDecimalPositive(data['payment'])) 
			errors.push({name: 'payment', message: 'Invalid amount!'});

		var payment = '0';
		if (data['payment'] && validateDecimalPositive(data['payment']))
			payment = data['payment'];
		payment = new Big(decimalToGwei(payment));
		
		var minPayment = $form.find('input[name="payment"]').attr('min');
		if (!minPayment)
			minPayment = '0';
		minPayment = new Big(minPayment);

		if (payment.lt(minPayment)) 
			errors.push({name: 'payment', message: 'Insufficient amount!'});

		return errors;
	}

	var transform = function (data) {
		
		var t = {};
		for (var k in data) 
			t[k] = data[k];
		
		var seedKey = t['seedKey'];
		delete t['seedKey'];
		t['seedKeyHash'] = h2d(sha256Hex(seedKey));
		
		var transformURI = function (name) {
			var value = t[name];
			if (validateCID(value)) {
				t[name] = 'ipfs://' + value;
			}
		}

		transformURI('codeURI');
		transformURI('imageURI');
		transformURI('projectURI');
		t['codeHash'] = h2d(t['codeHash']);

		if (t['payment']) 
			t['value'] = decimalToGwei(t['payment']);

		delete t['payment'];
		return t;
	}

	var contentCtx = ['div', {class: 'content'}, 

		walletConnector.ctx('mint'),

		['form', {class: 'form'},

			['div', {class: 'note'}, '* - Required fields'], 

			genFieldCtx('Name', 'name', 1, 
				'This is the name of your NFT. Don\'t worry about collisions.'), 
			
			genFieldCtx('Seed Key', 'seedKey', 1, 
				'This makes the off-chain RNG unpredictable. <br>'+
				'Only the hash will be stored on minting. <br>' +
				'After you have generated and released the results, you will need '+
				'to reveal this seed key for others to verify fairness. <br>' +
				'Make sure it is very strong and record it in a safe place.'), 
			
			genFieldCtx('Code Hash', 'codeHash', 1, 
				'This is the SHA-256 hash of the code zip in hexadecimal format.'), 
			
			genFieldCtx('Code URI', 'codeURI', 1, 
				'This is a permalink to the code zip. <br>' +
				'For best permanace, IPFS is recommend. <br>' + 
				'You can use web3.storage to store the files for free.'), 
			
			genFieldCtx('Image URI', 'imageURI', 0, 
				'An image for your NFT. You can change this later. <br>' +
				'If no image is provided, a random image will be generated on-chain.'), 
			
			genFieldCtx('Project URI', 'projectURI', 0,
				'Your project\'s URI. You can change this later. <br>' +
				'This can be an etherscan url or simply your project\'s website.'), 
			
			genFieldCtx(paymentLabel, 'payment', 0, 
				'A nominal fee to prevent spam and fund development. <br>' +
				'If you feel generous, you can pay more <br>' +
				'so that your project will appear above others in the search.'), 

			['div', {class: 'notice'},
				'Minting requires an additional ', ['span', {name: 'vrfFee'}, '0.0'], ' LINK ', 
				'for Chainlink VRF payment. <br>'
			],

			['button', {class: 'authorize green'}, authorizeLinkText], 

			['button', {class: 'mint'}, 'Mint'], 

			['div', {class: 'contract-link'}]
		]
	];

	var showContractLink = function () {
		var $c = $page.find('.contract-link');
		if (W.isSignedIn()) {
			var l = W.getEtherscanLink('chainrand');
			var m = l.match(/\/([0-9a-fA-Fx]+)$/);
			if (m) {
				$c.html(HTML(
					['a', {href: l, target: '_blank'}, 
					HTML.escape('View contract (' + m[1] + ')')]
				))	
			} else {
				$c.html('');	
			}
		} else {
			$c.html('');
		}
	}

	var init = function (sel) {
		$page = $(sel).html(HTML(contentCtx));
		$page.find('.help').click(function () {
			var t = $(this);
			var name = t.attr('name');
			var d = $(sel+' .description[name="' + name + '"]');
			
			if (t.hasClass('active')) {
				d.hide();
				t.removeClass('active');	
			} else {
				d.show();
				t.addClass('active');	
			}
		});
		$page.find('form').submit(function (e) {
			e.preventDefault();
		});
		$authorize = $page.find('.authorize');
		$mint = $page.find('.mint');
		disableElement($authorize);
		disableElement($mint);
		$form = $page.find('form');
		$form.find('input').each(function () {
			setupOnChange($(this).attr('name'));
		})
		walletConnector.sync('mint');
	}

	var resetPaymentField = function () {
		var $p = $form.find('input[name="payment"]');
		var m = $p.attr('min');
		var $m = $form.find('.f-row[name="payment"] .min');
		if (m != '0') {
			var v = gweiToDecimal(m);	
			$p.val(v);
			$m.show().find('.amount').html(v);
		} else {
			$m.hide();
			$p.val('');
		}
	}

	var clearForm = function () {
		$form.find('.error').hide();
		$form.find('input').val('');
		$form.find('.help').removeClass('active');
		$form.find('.description').hide();
		resetPaymentField();
	}

	var testPopulate = function () {
		$form.find('input[name="name"]')
			.val("Cr Test");
		$form.find('input[name="seedKey"]')
			.val("ThisIsAVeryLongKey");
		$form.find('input[name="codeHash"]')
			.val("0x5ef8642fa8d3ef3729d545c46039a46a0151528d3dd2cfbc2cd3763ce5cd1973");
		$form.find('input[name="codeURI"]')
			.val("https://sdw.mypinata.cloud/ipfs/QmUVQq4cvMcMVbdgRCeNzZEtF1EU4JmyqzoSdycxUAqNTz");
	}

	var waitForRandomness = function () {

		var consecDone = 0, prevId = -1;
		var recurse = function () {
			if (W.isSignedIn()) {
				W.callContract('chainrand', 'balanceOf', W.walletAddress, function (c) {
					console.log(c);
					c = parseInt(c, 10) - 1;
					if (c >= 0 && W.isSignedIn())
						W.callContract('chainrand', 'tokens', c, function (t) {
							console.log(t);
							var r = t.randomness + '';
							if (r != '0') {
								consecDone++;
								if (prevId != c) {
									tokensSyncCounter++;
									Layout.sync();
									prevId = c;
								}
							} 
							if (consecDone < 2)
								setTimeout(recurse, 1000 * 30);
						})
				})	
			}
		}
		setTimeout(recurse, 1000 * 15);
	}

	var onMinted = function (txHash) {
		clearForm();
		modal.alert('Mint submitted. <br>The tokens page will auto-refresh.');
		window.location.hash = '#/tokens';
		 
		setTimeout(function () { 
			W.waitTransactionMined(txHash, function () {
				console.log("TX Done!");
				tokensSyncCounter++;
				waitForRandomness();
				Layout.sync();
			})
		}, 30);
		console.log(txHash);
	}
	
	window._onMinted = onMinted;

	var mint = function (data) {
		var t = transform(data);
		console.log(t);
		W.mint(t.name, t.seedKeyHash, t.codeHash, t.codeURI, t.imageURI, t.projectURI, t.value, {
			success: onMinted, 
			error: function (e) {
				console.log(e)
			}
		})
	}

	var validateAndMint = function () {
		var data = toObject($form.serializeArray());
		var errors = validate(data);
		for (var i = errors.length - 1; i >= 0; i--) {
			var e = errors[i];
			var $r = $form.find('.f-row[name="' + e.name + '"]');
			$r.find('.error').html(e.message).show();
		}
		if (errors.length < 1) {
			mint(data);
		}
	}

	var enableMint = function () {
		$authorize.hide();
		enableElement($mint.show());
		$mint.off().click(validateAndMint);
	}

	var disableMint = function () {
		enableElement($authorize.show()
			.off().click(authorizeLink)
			.html(authorizeLinkText));
		disableElement($mint.show())
	}

	var onAuthorizedLinkChecked = function (hasAuthorized) {
		if (hasAuthorized) {
			enableMint();
		} else {
			disableMint();
		}
	}

	var authorizeLink = function () {
		W.authorizeLink({
			success: onAuthorizedLinkChecked, 
			pending: function () {
				disableElement($authorize
					.html('Waiting for authorization...'));
			},
			error: function (e) {
				console.log(e);
				disableMint();
			}
		})
	}

	var sync = function () {
		console.log(setTimeout(testPopulate, 300));
		walletConnector.sync('mint');
		if (W.chainId) {
			W.callContract('chainrand', 'vrfFee', function (result) {
				$page.find('[name="vrfFee"]').html(gweiToDecimal(result));
			});
			
			for (var i = 0; i < chainlist.length; ++i)
				if (chainlist[i].id == W.chainId)
					$page.find('.currency').html(chainlist[i].currency);
			
		}
		if (W.isSignedIn()) {
			var results = {}, done = 0;
			var setupMint = function () {
				if (done == 2) {
					// results.mintFee = '10000000000000000';
					$form.find('input[name="payment"]').attr('min', results.mintFee);
					resetPaymentField();
					onAuthorizedLinkChecked(results.hasAuthorized);
				}
			}
			W.callContract('chainrand', 'mintFee', function (mintFee) {
				results.mintFee = mintFee;
				done++;
				setupMint();
			})
			W.checkHasAuthorizedLink(function (hasAuthorized) {
				results.hasAuthorized = hasAuthorized;
				done++;
				setupMint();
			}); 
		}
		showContractLink();
	}

	Layout.add({
		label: 'mint', 
		init: init, 
		sync: sync,
	});
})()