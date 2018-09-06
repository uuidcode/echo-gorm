var core = {
    onClick: function (option) {
        $('body').on('click', option.buttonSelector, option.callback);
    },

    stringify: function (object) {
        return JSON.stringify(object, null, 4);
    },

    showModal: function (option) {
        $.get(option.url, function (html) {
            $('body').append(html);
            $(option.modalSelector).modal()
                .on('hide.bs.modal', function () {
                    $(this).remove();
                });
        })
    },

    initModal: function (option) {
        option.callback = function () {
            core.showModal(option)
        };

        core.onClick(option);
    },

    initSend: function (option) {
        option.callback = function () {
            $.ajax({
                url: option.url,
                type: 'POST',
                contentType: 'application/json; charset=utf-8',
                data: core.stringify(option.data()),
                success: function (json) {
                    location.reload();
                }
            });
        };

        core.onClick(option);
    }
};