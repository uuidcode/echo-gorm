var book = {
    init: function () {
        core.initModal({
            buttonSelector: '#addButton',
            modalSelector: '#bookModal',
            url: '/book/form'
        });

        core.initSend({
            buttonSelector: '#saveButton',
            url: '/book',
            data: function () {
                return {
                    name: $('#name').val()
                }
            }
        });
    }
};

$(function() {
    book.init();
});