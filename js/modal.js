var modal = (function () {
    
    var modalId = 'modal' + (Math.random() + '').substr(3);
    
    var self = function (p) {
        var $modal = $('#'+modalId);
        if ($modal.length < 1) {   
            var SVG = HTML.SVG, P = SVG.Path;
            var w = 18;
            var crossCtx = SVG(w,w, 
                P.M(0,0).L(w,w),
                P.M(0,w).L(w,0)
            );
            var modalCtx = ['div', {class: 'modal', id: modalId}, 
                ['div', {class: 'modal-dialog'},
                    ['div', {class: 'modal-content'}, 
                        ['div', {class: 'modal-header'},
                            ['div', {class: 'modal-close'}, 
                                crossCtx
                            ],
                            ['div', {class: 'modal-title'}]
                        ],
                        ['div', {class: 'modal-body'}],
                        ['div', {class: 'modal-footer'}]
                    ], 
                ],
                ['div', {class: 'modal-overlay'}]        
            ];
            $('body').append(HTML(modalCtx));
            $modal = $('#'+modalId);
        }

        if ('title' in p) $modal.find('.modal-title').html(p['title']).show();
        else $modal.find('.modal-title').hide();              
        
        if ('body' in p) $modal.find('.modal-body').html(p['body']).show();
        else $modal.find('.modal-body').hide();              

        if ('footer' in p) $modal.find('.modal-footer').html(p['footer']).show();
        else $modal.find('.modal-footer').hide();              

        self.close = function () {
            self.active = 0;
            $modal.removeClass('active');
            $('body').removeClass('modal-shown-body');
            setTimeout(function () {
                if (!self.active) $modal.hide()
            }, 300);
        };
        self.hide = self.close;
        
        self.active = 1;
        if ('beforeShow' in p && $.isFunction(p['beforeShow'])) {
            p['beforeShow']();
        }
        $modal.show();

        setTimeout(function () {
            $modal.addClass('active');
            $modal.find('.modal-close').click(self.close);
            $modal.find('.modal-overlay').click(self.close);
            if ('afterShow' in p && $.isFunction(p['afterShow'])) {
                setTimeout(p['afterShow'], 300);
            }
            $('body').addClass('modal-shown-body');
        }, 10);
    };

    var alertId = 'alert' + (Math.random() + '').substr(3);

    var modalAlert = function (p) {
        var $alert = $('#'+alertId);
        if ($alert.length < 1) {
            var alertCtx = ['div', {class: 'modal-alert', id: alertId}, 
                ['div', {class: 'modal-dialog'},
                    ['div', {class: 'modal-content'}, 
                        ['div', {class: 'modal-body'}],
                        ['div', {class: 'modal-footer'}]
                    ], 
                ],
                ['div', {class: 'modal-overlay'}]        
            ];
            $('body').append(HTML(alertCtx));
            $alert = $('#'+alertId);
        }         
        
        if ('body' in p) $alert.find('.modal-body').html(p['body']).show();
        else $alert.find('.modal-body').hide();              

        if ('footer' in p) $alert.find('.modal-footer').html(p['footer']).show();
        else $alert.find('.modal-footer').hide();              

        modalAlert.close = function () {
            modalAlert.active = 0;
            $alert.removeClass('active');
            $('body').removeClass('modal-alert-shown-body');
            setTimeout(function () {
                if (!modalAlert.active) $alert.hide()
            }, 300);
        };
        modalAlert.hide = modalAlert.close;
        
        modalAlert.active = 1;
        if ('beforeShow' in p && $.isFunction(p['beforeShow'])) {
            p['beforeShow']();
        }
        $alert.show();

        setTimeout(function () {
            $alert.addClass('active');
            $alert.find('.modal-close').click(modalAlert.close);
            $alert.find('.modal-overlay').click(modalAlert.close);
            if ('afterShow' in p && $.isFunction(p['afterShow'])) {
                setTimeout(p['afterShow'], 300);
            }
            $('body').addClass('modal-alert-shown-body');
        }, 10);
    };
    
    self.alert = modalAlert;

    return self;
})();