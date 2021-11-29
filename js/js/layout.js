var Layout = (function() {

	var self = {}

	var pages = []

	var pagesCtx = ['div', {id: 'pages'}];

	var navLinksCtx = [];

	var normalized = function (s) { 
		return s.replace(/\s+/g, '-').toLowerCase() 
	};

	var OPENSEA_URL = 'https://opensea.io/collection/sorasdreamworld';
	var TWITTER_URL = 'https://twitter.com/Sorasdreamworld';
	var DISCORD_URL = 'https://discord.gg/yVu7uWzmmy';

	var socialOpenseaCtx = ['a', {href: OPENSEA_URL, target: '_blank'}, ['img', {src: 'assets/s0cial_opensea.svg'}]];
	var socialTwitterCtx = ['a', {href: TWITTER_URL, target: '_blank'}, ['img', {src: 'assets/s0cial_tw1tter.svg'}]];
	var socialDiscordCtx = ['a', {href: DISCORD_URL, target: '_blank'}, ['img', {src: 'assets/s0cial_disc0rd.svg'}]];

	var socialTableCtx =  ['table', {class: 's0cial'}, ['tr', 
		['td', socialOpenseaCtx], 
		['td', socialTwitterCtx], 
		['td', socialDiscordCtx]
	]];

	socialTableCtx = ''; // We don't need it for now.

    var tryExec = function (f, k) {
        var fn = f ? f[k] : null, args = [];
        for (var i = 2; i < arguments.length; ++i) 
            args.push(arguments[i]);
        if ($.isFunction(fn)) {
            fn.apply(undefined, args);
        }
    };

    var rootSynced = false;

    var setupSlider = function () {
    	var slideAnimating = 0;
		$('#navbar .slider').hide();
		$('#navbar .hamburger').click(function () { 
			if (slideAnimating)
				return;
			slideAnimating = 1;
			var t = $(this), slider = $('#navbar .slider'), sliderLinks = slider.find('a');
			if (t.hasClass('active')) {
				setTimeout(function () {
					$('#navbar').removeClass('slided');
				}, 200);
				t.removeClass('active');
				slider.removeClass('active');
				sliderLinks.each(function(i, v) {
					setTimeout(function () {
						v.style.opacity = 0;
					}, Math.max(0, (sliderLinks.length - 3 - i) * 30));
				});
				setTimeout(function () { 
					slideAnimating = 0;
					slider.hide();
				}, 700);
			} else {
				slider.show();
				setTimeout(function () {
					t.addClass('active');
					slider.addClass('active');
					sliderLinks.each(function(i, v) {
						setTimeout(function () {
							v.style.opacity = 1;
						}, i * 60);
					});
					setTimeout(function () { slideAnimating = 0 }, 500);    
				}, 15);
				$('#navbar').addClass('slided');
			}
		});
		var hideMenu = function () {
			if (!slideAnimating) {
				$('#navbar .hamburger').click();
			} else {
				setTimeout(hideMenu, 60);
			}
		};
		$('#navbar .slider a, #navbar .slider .overlay').click(function () {
			var h = $(this).attr('href'), r = $(this).attr('target');
			if (h && (''+h).substr(0, 1) != '#' && r != '_blank') {  }
			else hideMenu();
		});
    }

	self.add = function (s) {
		
		s.path = normalized(s.label);
		s.synced = false;
		s.inited = false;
		s.visible = false;
		s.sel = '#pages .page[path="' + s.path + '"]';

		pages.push(s);
		
		pagesCtx.push(['div', {class: 'page', path: s.path}]);

		navLinksCtx.push(['a', {href: '#/' + s.path}, capitalize(s.label)]);

		var onRoute = function () {
			for (var i = 0; i < pages.length; ++i) {
				var p = pages[i];
				if (p.path == s.path) {
					p.visible = true;
					if (!p.inited) {
						tryExec(p, 'init', p.sel);
						p.inited = true;
					}
					$(p.sel).show();
					tryExec(p, 'show', p.sel);		
					if (rootSynced) {
						if (!p.synced) {
							tryExec(p, 'sync', p.sel);
							p.synced = true;
						}	
					}
				} else {
					$(p.sel).hide();
					tryExec(p, 'hide', p.sel);
					p.visible = false;
				}
			}
		}

		HTML.route(s.path, onRoute);
	}

	self.init = function () {	

		var navCtx = ['div', {id: 'navbar'},
			['div', {class: 'desktop'}, ['table', ['tr',
				['td', {class: 'c-logo'}, ['a', {href:'#/home'}, ['img', {src: 'assets/nav_logo.png'}]]],
				['td', {class: 'c-links'}, navLinksCtx],
				['td', {class: 'c-s0cial'}, socialTableCtx]
			]]],
			['div', {class: 'mobile'}, ['table',  ['tr',
				['td', {class: 'c-hamburger'}, ['div', {class: 'hamburger'},
					['div', {class: 'l-0'}],
					['div', {class: 'l-1'}],
					['div', {class: 'l-2'}],
				]],
				['td', {class: 'c-logo'}, ['img', {src: 'assets/nav_logo.png'}]],
				['td', {class: 'c-right'}],
			]]], 
			['div', {class: 'slider'}, 
				['div', {class: 'c-links'}, navLinksCtx], 
				['div', {class: 'c-s0cial'}, socialTableCtx],
				['div', {class: 'bg'}],
				['div', {class: 'overlay'}],
			]
		];

		$('body').html(HTML([navCtx, pagesCtx]));

		HTML.route.go(':home');

		setupSlider();
	}


	self.sync = function () {
		rootSynced = true;
		for (var i = 0; i < pages.length; ++i) {
			var p = pages[i];
			if (p.visible) {
				tryExec(p, 'sync', p.sel)
				p.synced = true;
			} else {
				p.synced = false;
			}
		}
	}

	return self;

})();