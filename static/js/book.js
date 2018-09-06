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

        $('body').on('click', '#saveButton', function () {
            $.ajax({
                url: '/book',
                type: 'POST',
                contentType: 'application/json; charset=utf-8',
                data: book.stringify({
                    name: $('#name').val()
                }),
                success: function (json) {
                    location.reload();
                }
            });
        });
    },

    stringify: function (object) {
        return JSON.stringify(object, null, 4);
    }
};

$(function() {
    book.init();
});