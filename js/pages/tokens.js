(function () {

	var $page, $window, $body;

	var contentCtx = ['div', {class: 'content'},
		
		walletConnector.ctx('tokens'),

		['div', {class: 'results'}]
	];

	var visible = 0, retrievingTokens = 0, batchSize = 5, displayed = '';

	var displayPageMessage = function (message) {
		var ctx = ['div', {class: 'message'}, message];
		$page.find('.results').html(HTML(ctx));
	}

	var edit = function () {
		var t = $(this);
		var tokenId = t.attr('tokenId');
		var name = t.attr('name');
		var label = t.attr('label');
		var title = 'Change ' + label;
		var saveText = 'Save';
		if (name == 'seedKey') {
			title = 'Reveal Seed Key';
			saveText = 'Reveal';
		}
		if (name == 'paid') {
			title = 'Pay';
			var currency = '';
			for (var i = 0; i < chainlist.length; ++i)
				if (chainlist[i].id == W.chainId)
					currency = chainlist[i].currency;
			label = 'Amount (' + currency + ')';
			saveText = 'Pay';
		}
		
		var bodyCtx = ['form', {class: 'edit-body'}, ['table', ['tr',
			['td', {class: 'label'}, ['span', label + ': ']], 
			['td', {class: 'value'}, 
				['input', {type: 'text', name: name}]
			]
		]]];
		var footerCtx = ['div', {class: 'edit-footer'},
			['button', {class: 'light cancel'}, 'Cancel'], 
			['div', {class: 'spacer'}],
			['button', {class: 'save'}, saveText] 
		];
		var beforeShow = function () {
			setTimeout(function () {
				var $label = $('.edit-body .label');
				var w = $label.find('span').width() + 18;
				$label.css('width', w);
				$('.edit-body .error').hide();
			}, 60)
		};
		
		var submit = function () {
			var data = toObject($('.edit-body').serializeArray()); 
			var value = data[name];
			var error = '';
			var isEmpty = !value;
			var confirmMsg = '';
			console.log(value + ' ' + name)
			if (name == 'paid') {
				if (!isEmpty && !validateDecimalPositive(value))
					error = 'Invalid amount!';
				value = isEmpty ? '0' : value;
				if (validateDecimalPositive(value))
					value = decimalToGwei(value);
				if (value == '0')
					confirmMsg = 'Are you sure you want to pay nothing? <br>' +
						'You will still incur gas fees';
			}
			if (name.substr(-3) == 'URI') {
				if (isEmpty) {
					confirmMsg = 'Are you sure you want to submit an empty value?<br>' +
					'This will remove the current value.';
				} else {
					if (!validateURL(value) && !validateCID(value))
						error = 'Must be a valid URI or CID!';
					else if (validateCID(value))
						value = 'ipfs://' + value;
				}
			}
			var fns = {
				success: function (txHash) {
					console.log(txHash);
					W.waitTransactionMined(txHash, function (status) {
						console.log(status);
						displayed = '';
						scrollToTop();
						setTimeout(function () { Layout.sync() }, 600);
					})
					modal.close();
				}, 
				error: function (e) {
					console.log(e);
					if (e.message) {
						modal.alert(e.message);
					} else {
						modal.alert('An error occured. Please reload page.');
					}
				}
			};
			var execute = function () {
				if (name == 'paid') {
					W.payToken(tokenId, value, fns);
				} else {
					W.updateToken(tokenId, name, value, fns);
				}
			};
			if (error) {
				modal.alert(error);
				return
			} 
			if (confirmMsg) {
				modal.confirm(confirmMsg, execute);
			} else {
				execute();
			}
		};

		var afterShow = function () {
			var $label = $('.edit-body').submit(function (e) {
				e.preventDefault();
				submit();
			});
			var $footer = $('.edit-footer');
			$footer.find('.cancel').click(modal.close);
			$footer.find('.save').click(submit);
		}

		modal({
			title: title, 
			body: HTML(bodyCtx), 
			footer: HTML(footerCtx), 
			beforeShow: beforeShow,
			afterShow: afterShow
		})
	}

	var dataFields = 'id,name,codeURI,seedKey,imageURI,projectURI,minter,seedKeyHash,codeHash,randomness,paid,verified'.split(',');

	var genTokenCtx = function (t, chainId) {
		
		t.verified = t.verified == 1 ? 'Yes' : 'No';
		t.paid = gweiToDecimal(t.paid);
		t.openSeaURI = genOpenSeaURL(t.id, chainId);

		var rowsCtx = ['table']; 
		
		var pushRowCtx = function (label, editIconClass) {
			var name = label.replace(/\s/g, '');
			name = name.substr(0, 1).toLowerCase() + name.substr(1);
			var value = t[name];
			var isURI = label.substr(-3) == 'URI';
			if (value || editIconClass) {
				var e = HTML.escape(value);
				if (isURI && value) 
					e = ['a', {href: value, target: '_blank'}, e];
				e = ['div', {class: 'code'}, e];
				if (editIconClass) {
					if (!value)
						e.push(['span', {class: 'empty'}, HTML.escape('<Empty>')]);	
					var iconCtx = ['i', {class: editIconClass}];
					e.push({class: 'editable'}, 
						['div', {class: 'edit', tokenId: t.id, name: name, label: label}, iconCtx]);
				}
				rowsCtx.push(['tr',
					['td', {class: 'label'}, label + ': '],
					['td', {class: 'value'}, e],
				]);	
			}
		};

		rowsCtx.push(['tr', ['td', {colspan: 2},
			['div', {class: 'head'},
				['span', {class: 'id'}, '#' + t.id], 
				['span', {class: 'name'}, HTML.escape(t.name)]
			]
		]]);

		pushRowCtx('Verified');
		pushRowCtx('Project URI', 'icon-pencil');
		pushRowCtx('Image URI', 'icon-pencil');
		pushRowCtx('Code URI', 'icon-pencil');
		pushRowCtx('Code Hash');
		pushRowCtx('Seed Key', 'icon-key');
		pushRowCtx('Seed Key Hash');
		pushRowCtx('Paid', 'icon-plus');
		pushRowCtx('Randomness');
		pushRowCtx('OpenSea URI');

		rowsCtx.push(['tr', ['td', {colspan: 2},
			['div', {class: 'foot'},
				['a', {href: t.openSeaURI, target: '_blank'}, 
					'View on OpenSea', ['i', {class: 'icon-link-ext'}]]
			]
		]])
		
		return ['div', {class: 'token'}, rowsCtx]
	}

	var retrieveNextBatch = function () {
		if (retrievingTokens) return;	
		var $next = $page.find('.batch[done=0]').first();
		if ($next.length > 0) {
			retrievingTokens = 1;
			var tokenIds = $next.attr('batchIds').split(',');
			W.callContract('chainrand', 
				'tokenData', tokenIds, 
				function (data) {
					console.log(data);
					retrievingTokens = 0;
					$next.attr('done', '1');
					var chainId = $next.attr('chainId');
					var step = dataFields.length;
					var ctx = ['div'];
					for (var i = 0; i < data.length; i += step) {
						var t = {};
						for (var j = 0; j < step; ++j)
							t[dataFields[j]] = data[i+j];
						ctx.push(genTokenCtx(t, chainId));
					}
					$next.html(HTML(ctx)).find('.edit').click(edit);
				});
		}
	}

	var onScroll = function () {
		if (visible) {
			var scrollTop = $window.scrollTop();
			var scollableDist = $body.height() - $window.height();
			if (scollableDist - scrollTop < 500) {
				retrieveNextBatch();
			}
		}
	}

	var init = function (sel) {
		$page = $(sel).html(HTML(contentCtx));
		walletConnector.sync('tokens');
		$window = $(window); 
		$body = $('body');
		$window.scroll(throttle(onScroll, 300));
	}

	var displayTokenIds = function (tokenIds, walletAddress, chainId) {
		if (tokenIds.length < 1) {
			displayPageMessage('Wallet has no tokens.');
			return
		}

		var currDisplayed = tokenIds.join(',') + '/' + walletAddress + '/' + chainId;
		if (currDisplayed == displayed) return;

		var batchesCtx = ['div', {class: 'batches'}];
		var n = tokenIds.length;
		for (var i = 0; i < n; i += batchSize) {
			var batchIds = [];
			for (var j = 0; j < batchSize && i+j < n; ++j) {
				batchIds.push(tokenIds[(n-1)-(i+j)]);
			}
			batchesCtx.push(['div', {
				class: 'batch', 
				batchIds: batchIds.join(','), 
				id: elId(), 
				walletAddress: walletAddress, 
				chainId: chainId,
				done: '0'
			}]);
		}
		displayed = currDisplayed;
		$page.find('.results').html(HTML(batchesCtx));
		retrieveNextBatch();
	}
	window._fx = function () {
		window._onMinted('0x7e05cc3ad782da6cf5aa38195d8f91080e6f967f4799dc5a2c2098daef41b3ac');
	}
	window._ret = retrieveNextBatch;
	var sync = function (sel) {
		console.log(W.walletAddress + ' ' + W.chainId);
		if (W.isSignedIn()) {
			displayPageMessage('Loading...');
			W.callContract('chainrand', 'tokensOfOwner', W.walletAddress, 
				function (tokenIds) {
					console.log(tokenIds);
					displayTokenIds(tokenIds, W.walletAddress, W.chainId);
				});	
		} else {
			displayPageMessage('Not connected.')
		}
		walletConnector.sync('tokens');
	}

	var show = function () {
		visible = 1;
	}

	var hide = function () {
		visible = 0;
	}

	Layout.add({
		label: 'tokens', 
		init: init, 
		sync: sync,
		show: show, 
		hide: hide
	});
})()