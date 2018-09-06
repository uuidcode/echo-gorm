var book = {
    init: function () {
        $('#addButton').on('click', function () {
            $.get('/book/form', function (html) {
                $('body').append(html);
                $('#bookModal').modal()
                    .on('hide.bs.modal', function () {
                        $(this).remove();
                    });
            })
        })
    }
};

$(function() {
    book.init();
});