(function () {

	var $page, $form;
	
	var contextCtx = ['div', {class: 'content'},
		['form', {class: 'form'},
			['div', {class: 'params'}, 
				['div', {class: 'type'},
					['name', 'minter'], function (value) {
						var id = elId();
						return ['span', {class: 'radio'},
							['input', {type: 'radio', id: id, value: value, name: 'type'}],
							['div', {class: 'check'}, ['div']],
							['label', {for: id}, capitalize(value)]
						]					
					}
				],
				['div', {class: 'chain'}, 
					['select', {name: 'chain'}, 
						chainlist, function (c) { 
							return c.supported ? ['option', {value: c.id}, capitalize(c.name)] : '';
						}
					]
				]
			],
			['div', {class: 'query'},
				['input', {type: 'text', name: 'query'}], 
			], 
			['div', {class: 'search'},
				['button', 'Search'], 
				['input', {type: 'submit'}]
			]
		], 
		['div', {class: 'results'}]
	];

	var search = function (data) {
		var h = '#/search/' + data.chain + '/' + data.type + '/' + 
			encodeURIComponent(data.query);
		window.location.hash = h;
	}

	var init = function (sel) {
		$page = $(sel).html(HTML(contextCtx));
		$page.find('.type input[type="radio"]').first().prop('checked', true);
		$form = $page.find('form').submit(function (e) {
			e.preventDefault();
			var t = $(this);
			search(toObject(t.serializeArray()));
		});
	}

	var sync = function () {
		if (W.chainId) {
			$page.find('.chain select').val(W.chainId)
		}
	}

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
						['div', {class: 'edit', tokenId: t.id}, iconCtx]);
				}
				rowsCtx.push(['tr', 
					['td', {class: 'label'}, label + ': '],
					['td', {class: 'value', name: name}, e],
				]);	
			}
		};

		rowsCtx.push(['tr', ['td', {colspan: 2},
			['div', {class: 'head'},
				['span', {class: 'id'}, '#' + t.id], 
				['span', {class: 'name'}, HTML.escape(t.name)]
			]
		]]);

		pushRowCtx('Minter');
		pushRowCtx('Verified');
		pushRowCtx('Project URI');
		pushRowCtx('Image URI');
		pushRowCtx('Code URI');
		pushRowCtx('Code Hash');
		pushRowCtx('Seed Key');
		pushRowCtx('Seed Key Hash');
		pushRowCtx('Randomness');

		rowsCtx.push(['tr', ['td', {colspan: 2},
			['div', {class: 'foot'},
				['a', {href: t.openSeaURI, target: '_blank'}, 
					'View on OpenSea', ['i', {class: 'icon-link-ext'}]]
			]
		]])
		
		return ['div', {class: 'token'}, rowsCtx]
	}

	var displayNoResults = function () {
		$page.find('.results')
			.html(HTML(['div', {class: 'message'}, 'No results.']));
	}

	var displayNonce = 0;

	var displayTokens = function (tokens, chainId) {
		var nTriesLeft = 5;
		var d = ++displayNonce;
		var waitSync = function () {
			if (W.config) {
				var tempId = elId();
				var resultsCtx = ['div', tokens, function (t) {
					return genTokenCtx(t, chainId)
				}, {id: tempId}];
				if (d == displayNonce) {
					$page.find('.results').html(HTML(resultsCtx));
				}
				addTokenTippys(tempId);
			} else if (nTriesLeft > 0) {
				nTriesLeft--;
				setTimeout(waitSync, 500);
			} else {
				displayNoResults();
			}
		}
		waitSync();
	}

	var show = function () {
		var r = window.location.hash;
		var m = r.match(/^(?:#\/)?search\/(\d+)\/([^\/]+)\/(.*)$/i);
		if (m) {
			var data = {};
			var query = decodeURIComponent(m[3]);;
			var type = m[2].toLowerCase();
			var chainId = m[1];
			data[type] = query;
			data['c'] = chainId;
			$form.find('input[name="query"]').val(query);
			$form.find('select[name="chain"]').val(chainId);
			$form.find('input[name="type"]').each(function () {
				if ($(this).attr('value') == type) 
					$(this).prop('checked', type);
			})
			$.ajax({
				url: 'api',
				type: 'GET',
				data: data, 
				success: function (r) {
					if (r.tokens && r.tokens.length > 0) {
						displayTokens(r.tokens, chainId)
					} else {
						displayNoResults();
					}
				}
			})
		}
	}

	Layout.add({
		label: 'search', 
		init: init, 
		sync: sync,
		show: show
	});
})()