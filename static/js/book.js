var book = {
    init: function () {
        core.initModal({
            buttonSelector: '#addButton',
            modalSelector: '#bookModal',
            url: '/book/form'
        });

        core.initAction({
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