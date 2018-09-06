var core = {
    onClick: function (option) {
        $('body').on('click', option.buttonSelector, function () {
            option.callback($(this));
        });
    },

    stringify: function (object) {
        return JSON.stringify(object, null, 4);
    },

    showModal: function (option) {
        var url = option.url || option.$target.attr('data-url');
        console.log('url', url);

        $.get(url, function (html) {
            $(html).appendTo('body')
                .modal()
                .on('hide.bs.modal', function () {
                    $(this).remove();
                });
        })
    },

    /**
     * buttonSelector: modal를 띄우는 버튼
     * url: modal html를 출력하는 url
     *  url이 없으면 button.attr('data-url')를 url로 사용합니다.
     */
    initModal: function (option) {
        option = option || {
            buttonSelector: '.openModalButton'
        };

        option.callback = function ($target) {
            option.$target = $target;
            core.showModal(option)
        };

        core.onClick(option);
    },

    /**
     * buttonSelector: 데이터를 전송송하는 버튼
     * data: 저장할 데이터는 가져오는 함수
     *   data가 없으면 button.attr('data-function')를 함수로 사용합니다.
     *   button.attr('data-function')가 없으며 해당 버튼을 포함하고 있는 modal의 form 요소를 데이터를 생성합니다.
     * url: 데이터를 저장하는 url
     *  url이 없으면 button.attr('data-url')를 url로 사용합니다.
     */
    initAction: function (option) {
        option = option || {
            buttonSelector: '.actionButton'
        };

        option.callback = function ($target) {
            var url = option.url || $target.attr('data-url');
            var data = {};
            var type = $target.attr('data-type') || 'POST';

            if (type == 'DELETE') {
                var id = $target.attr('data-id-name');
                data[id] = $target.attr('data-id');
            } else {
                if (option.data) {
                    data = option.data();
                } else {
                    var dataFunctionName = $target.attr('data-function')

                    if (dataFunctionName) {
                        data = dataFunctionName();
                    } else {
                        data = core.getData($target);
                    }
                }
            }

            $.ajax({
                url: url,
                type: type,
                contentType: 'application/json; charset=utf-8',
                data: core.stringify(data),
                success: function (json) {
                    location.reload();
                }
            });
        };

        core.onClick(option);
    },

    getData: function ($target) {
        var $modal = $target.closest('div.modal');
        var data = {};

        $modal.find('input,select,textarea')
            .map(function () {
                var id = $(this).attr('id');
                data[id] = $(this).val();
            });

        console.log('data', data);

        return data;
    },

    init: function () {
        core.initAction();
        core.initModal();
    }
};

$(function() {
    core.init();
});
