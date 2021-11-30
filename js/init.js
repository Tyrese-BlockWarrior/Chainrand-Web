$(function () {
		
    Layout.init();

    $.ajax({
        url: CONTRACTS_CONFIG_URL,
        success: function (e) {
            W.config = e;
            var s = function () { Layout.sync () };
            setupEth(W, {success: s, error: s});
        }
    });
})