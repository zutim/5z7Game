//数据表格
!(function ($) {

    var tableapi = layui.table;

    var storeInfoTpl = function (sysID, storeName, conName, conMobile, conPhoto, areaName) {

        var tplID = 'tpl-common-storeinfo' + sysID;
        var tpl = "<script type='text/html' id='" + tplID + "'>";
        tpl += "<div class='common-store-info'>"

        var gq = "{{d." + sysID + "==2?\"<span class='g'>供</span>\":\"<span class='q'>求</span>\"}}"

        tpl += "<img class='photo' src='{{(d." + conPhoto + "==''||d." + conPhoto + "=='null'||d." + conPhoto + "==null)?'res/common/images/003.jpg':d." + conPhoto + "}}' />";
        tpl += "<div class='store'>" + gq + "{{d." + storeName + "}}</div>";

        tpl += "<div class='address'>{{d." + areaName + "}}</div>"

        tpl += "<div class='attr'><span>{{d." + conName + "}}</span><span>{{d." + conMobile + "}}</span></div>";


        tpl += "</div>";
        tpl += "</script>";


        if ($("#" + tplID).length == 0) {
            $("body").append(tpl);
        }
        return '#' + tplID;
    };

    function _getOptions(dom) {

        _getSort(dom);
        _getWhere(dom);

        var op = {
            id: $(dom).attr('id') + '_tb',
            elem: '#' + $(dom).attr('id'),
            url: $(dom).attr('data-url'),
            method: 'post',
            cols: [_getCols(dom)],
            //skin: 'line',
            limit: 20,
            done: function (d) {
                $(dom).data('callback')(d);
                $(".common-store-info").click(function () {
                });
            },
            limits: [10, 20, 50, 100],
            text: {
                none: '暂无相关数据' //默认：无数据。注：该属性为 layui 2.2.5 开始新增
            },
            loading: true,
            where: $(dom).data('where')
        };

        if (typeof $(dom).attr("data-height") != "undefined")
            op.height = $(dom).attr('data-height');

        if (typeof $(dom).attr("size") != "undefined")
            op.size = $(dom).attr('size');

        if (typeof $(dom).attr("data-page") != "undefined")
            op.page = $(dom).attr('data-page') == "" ? true : $(dom).attr('data-page');

        if (!op.page)
            op.limit = 10000;

        var sort = $(dom).data('sort');
        if (typeof sort['type'] != 'undefined')
            op.initSort = sort;

        return op;
    }

    function _getCols(dom) {
        var cols = [];
        $(dom).find('thead').find('th').each(function (i, item) {
            var newCol = {};

            newCol.title = $(item).text();

            if (typeof $(item).attr("checkbox") != "undefined")
                newCol.type = 'checkbox';

            if (typeof $(item).attr("field") != "undefined")
                newCol.field = $(item).attr("field");

            if (typeof $(item).attr("width") != "undefined")
                newCol.width = $(item).attr("width");

            if (typeof $(item).attr("minWidth") != "undefined")
                newCol.minWidth = $(item).attr("minWidth");

            if (typeof $(item).attr("sort") != "undefined")
                newCol.sort = true;

            if (typeof $(item).attr("templet") != "undefined") {
                var tpl = $(item).attr("templet");
                if (tpl == 'store-common-info') {
                    var sysID = $(item).attr('sys-id');
                    var storeName = $(item).attr('store-name');
                    var contactName = $(item).attr('contact-name');
                    var contactMobile = $(item).attr('contact-mobile');
                    var contactPhoto = $(item).attr('contact-photo');
                    var areaName = $(item).attr('area-name');
                    newCol.templet = storeInfoTpl(sysID, storeName, contactName, contactMobile, contactPhoto, areaName);
                } else
                    newCol.templet = tpl;
            }

            if (typeof $(item).attr("toolbar") != "undefined")
                newCol.toolbar = $(item).attr("toolbar");

            if (typeof $(item).attr("align") != "undefined")
                newCol.align = $(item).attr("align");

            if (typeof $(item).attr("fixed") != "undefined")
                newCol.fixed = $(item).attr("fixed");

            if (typeof $(item).attr("event") != "undefined")
                newCol.event = $(item).attr("event");

            if (typeof $(item).attr("edit") != "undefined")
                newCol.edit = $(item).attr("edit");

            cols.push(newCol);
        });

        return cols;
    }

    function _getSort(dom) {

        var sort = $(dom).data('sort');
        if (typeof sort['type'] != 'undefined')
            return $(dom).data('sort');

        var data = {};
        var flag = true;
        $(dom).find('thead').find('th').each(function (i, item) {
            if (typeof $(item).attr("sort") != "undefined")
                if ($(item).attr("sort") != '' && flag) {
                    data = {field: $(item).attr("field"), type: $(item).attr("sort")};
                    flag = false;
                }
        });
        $(dom).data('sort', data);
    }

    function _getWhere(dom) {

        var boxObj = $(dom).attr("data-search");
        var data = {};
        $(boxObj).find('.search-key').each(function (i, item) {
            var key = $(item).attr('data-key');
            switch ($(item).attr('data-type')) {
                case 'txt':
                    var value = $(item).find('input').val();
                    if ($.trim(value) != '')
                        data[key] = value;
                    break;
                case 'list':
                    if ($(item).find('.cur').length > 0)
                        data[key] = $(item).find('.cur').find('span').attr('data-value');
                    break;
                case 'select':
                    var svalue = $(item).find('select').val();
                    if ($.trim(svalue) != '')
                        data[key] = svalue;
                    break;
            }
        });

        var sort = $(dom).data('sort');
        if (typeof sort['type'] != 'undefined') {
            if (typeof sort.type == 'string') {
                data.sortkey = sort.field;
                data.sorttype = sort.type;
            }
        }
        $(dom).data('where', data);
    }


    function _reload(dom) {
        var op = _getOptions(dom);
        tableapi.render(op);
        //tableapi.reload($(dom).attr('id') + '_tb',op);
    }

    function _eventSearch(dom) {
        if (typeof $(dom).attr("data-search") != "undefined") {
            var boxObj = $(dom).attr("data-search");

            $(boxObj).find('.updown > span').click(function () {
                if ($(this).parent().hasClass('more')) {
                    $(boxObj).find('.hide').show();
                    //$("#data-list").attr("data-height","full-210");
                }
                else
                    $(boxObj).find('.hide').hide();
                $(this).parent().toggleClass('more');
            });

            $(boxObj).find('.btn-query').click(function () {
                _reload(dom);
            });

            $(boxObj).find('.search-key').each(function (i, item) {

                switch ($(item).attr('data-type')) {
                    case 'txt':

                        break;
                    case 'list':
                        var fn = $(item).attr('data-callback') || null;
                        $(item).find('li').click(function () {

                            if ($(this).hasClass('cur'))
                                $(item).find('li').removeClass('cur');
                            else
                                $(this).siblings().removeClass('cur').end().addClass('cur');

                            if (fn != null)
                                eval(fn + '(this)');

                            _reload(dom);
                        });
                        break;
                    case "select":
                        $(item).find('select').change(function () {
                            _reload(dom);
                        });
                        break;
                }
            });

        }
    }


    function _eventSort(dom) {

        var filter = $(dom).attr('lay-filter');

        tableapi.on('sort(' + filter + ')', function (obj) {

            var sort = $(dom).data('sort');
            sort.field = obj.field;
            sort.type = obj.type;
            $(dom).data('sort', sort);

            _reload(dom);
        });
    }

    $.fn.table = function (callback) {
        callback = callback || function () {

        };
        var _that = this;
        $(this).data('where', {});
        $(this).data('sort', {});
        $(this).data('callback', callback);

        var op = _getOptions(this);
        // if (op.page) {
        //     var pdata = {};
        //     for (var k in op.where)
        //         pdata[k] = op.where[k];
        //     pdata.limit = 0;
        //     $.post(op.url, pdata, function (d) {
        //         if (parseInt(d.count) < op.limit) {
        //             op.height = "auto";
        //             op.page = false;
        //         }
        //         tableapi.render(op);
        //     }, "json");
        // } else
        tableapi.render(op);


        _eventSort(this);
        _eventSearch(this);


        return {
            search: function () {
                _reload(_that);
            }
        };
    }

})(jQuery);

//上传图片
!(function ($) {

    function preventBubble(event) {
        var e = arguments.callee.caller.arguments[0] || event; //若省略此句，下面的e改为event，IE运行可以，但是其他浏览器就不兼容
        if (e && e.stopPropagation) {
            e.stopPropagation();
        } else if (window.event) {
            window.event.cancelBubble = true;
        }
    }

    var upload = layui.upload;

    function _boxTemp() {
        var temp = '<div class="upload-images-box"><ul></ul></div>';
        return $(temp);
    }

    function _itemTemp(src) {
        src = src.indexOf("http") != -1 ? src : (resUrl + src);
        var temp = '<li class="item">' +
            '               <div>' +
            '                   <table>' +
            '                       <tr>' +
            '                           <td><img src="' + src + '" data-path="' + src + '"></td>' +
            '                       </tr>' +
            '                   </table>' +
            '                   <a target="_blank" href="' + src + '" class="ico view fa fa-search-plus"></a><i class="ico del fa fa-trash"></i>' +
            '               </div>' +
            '           </li>';
        return $(temp);
    }

    function _addTemp() {
        var temp = '<li class="add">' +
            '           <div>' +
            '               <i class="fa fa-file-photo-o"></i>' +
            '               <span>点击上传</span>' +
            '           </div>' +
            '       </li>';
        return $(temp);
    }

    function _initUpload(obj, dir, count, fn) {
        obj.on("click", function () {
            imageService(function (files) {

                files.map(function (f, i) {
                    var d = {data: {path: f.path, src: f.src}};
                    fn(d);
                });

            }, count > 1 ? true : false);
        });
    }

    function _refreshAddStatus(box, count) {
        if (box.find('.item').length >= count)
            $(box).find('.add').hide();
        else
            $(box).find('.add').show();
    }

    function _refreshSrc(obj, box) {
        var src = '';
        $(box).find('img').each(function () {
            if (src == '')
                src += $(this).attr('data-path');
            else
                src += "," + $(this).attr('data-path')
        });
        $(obj).val(src);
    }

    $.fn.uploadImages = function (flag) {
        flag = flag || false;
        $(this).each(function (i, inputObj) {

            var count = (typeof $(inputObj).attr("count") == "undefined") ? 1 : $(inputObj).attr("count");
            var dir = $(inputObj).attr("dir");

            var box = _boxTemp();

            var srcs = $(inputObj).val();
            var arrSrc = srcs.split(',');
            $(arrSrc).each(function (k, src) {
                if ($.trim(src) != '')
                    box.find('ul').append(_itemTemp(src));
            });

            box.find('ul').append(_addTemp());
            _refreshAddStatus(box, count);

            $(inputObj).parent().append(box);


            _initUpload($(box).find('.add > div'), dir, count, function (res) {
                var newItem = _itemTemp(flag ? res.data.src : res.data.path);

                $(box).find('.add').before(newItem);

                _initUpload(newItem, dir, 1, function (res) {
                    $(newItem).find('td').find('img').attr('src', res.data.src);
                    $(newItem).find('td').find('img').attr('data-path', res.data.src);
                    $(newItem).find('td').find('img').parents("table").next().attr("href", res.data.path);

                    _refreshAddStatus(box, count);
                    _refreshSrc(inputObj, box);
                });

                _refreshAddStatus(box, count);
                _refreshSrc(inputObj, box);

                $(newItem).find('.del').click(function () {
                    preventBubble();
                    $(newItem).remove();
                    _refreshAddStatus(box, count);
                    _refreshSrc(inputObj, box);
                });

                $(newItem).find('.view').click(function () {
                    preventBubble();
                });
            });

            $(box).find('.item').each(function (r, row) {
                _initUpload($(row).find('td'), dir, 1, function (res) {
                    $(row).find('td').find('img').attr('src', res.data.path);
                    $(row).find('td').find('img').attr("data-path", res.data.path);
                    $(row).find('td').find('img').parents("table").next().attr("href", res.data.path);
                    _refreshSrc(inputObj, box);
                    _refreshAddStatus(box, count);
                    console.log(22)
                });

                $(row).find('.del').click(function () {
                    preventBubble();
                    $(row).remove();
                    _refreshAddStatus(box, count);
                    _refreshSrc(inputObj, box);
                });

                $(row).find('.view').click(function () {
                    preventBubble();
                });
            })


        })
    }
})(jQuery);


//上传视频
!(function ($) {

    function preventBubble(event) {
        var e = arguments.callee.caller.arguments[0] || event; //若省略此句，下面的e改为event，IE运行可以，但是其他浏览器就不兼容
        if (e && e.stopPropagation) {
            e.stopPropagation();
        } else if (window.event) {
            window.event.cancelBubble = true;
        }
    }

    var upload = layui.upload;

    function _boxTemp() {
        var temp = '<div class="upload-images-box"><ul></ul></div>';
        return $(temp);
    }

    function _itemTemp(src) {
        var temp = '<li class="item">' +
            '               <div >' +
            '                           <div style="background-color: #000;width: 100px;height: 100px;overflow: hidden;"><video id="media" width="100" height="100" controls>' +
            '<source src="' + src + '">' +
            '</video></div>' +
            '                   <a target="_blank" href="' + src + '" class="ico view fa fa-search-plus"></a><i class="ico del fa fa-trash"></i>' +
            '               </div>' +
            '           </li>';
        return $(temp);
    }

    function _addTemp() {
        var temp = '<li class="add">' +
            '           <div>' +
            '               <i class="fa fa-file-movie-o"></i>' +
            '               <span>点击上传</span>' +
            '           </div>' +
            '       </li>';
        return $(temp);
    }

    function _initUpload(obj, dir, count, fn) {
        obj.on("click", function () {
            videoService(function (files) {

                files.map(function (f, i) {
                    var d = {data: {path: f.path, src: f.src}};
                    fn(d);
                });

            }, count > 1 ? true : false);
        });
    }

    function _refreshAddStatus(box, count) {
        if (box.find('.item').length >= count)
            $(box).find('.add').hide();
        else
            $(box).find('.add').show();
    }

    function _refreshSrc(obj, box) {
        var src = '';
        $(box).find('source').each(function () {
            if (src == '')
                src += $(this).attr('src');
            else
                src += "," + $(this).attr('src')
        });
        $(obj).val(src);
    }

    $.fn.uploadVideo = function (flag) {
        flag = flag || false;
        $(this).each(function (i, inputObj) {

            var count = (typeof $(inputObj).attr("count") == "undefined") ? 1 : $(inputObj).attr("count");
            var dir = $(inputObj).attr("dir");

            var box = _boxTemp();

            var srcs = $(inputObj).val();
            var arrSrc = srcs.split(',');
            $(arrSrc).each(function (k, src) {
                if ($.trim(src) != '')
                    box.find('ul').append(_itemTemp(src));
            });

            box.find('ul').append(_addTemp());
            _refreshAddStatus(box, count);

            $(inputObj).parent().append(box);


            _initUpload($(box).find('.add > div'), dir, count, function (res) {
                var newItem = _itemTemp(flag ? res.data.src : res.data.path);

                $(box).find('.add').before(newItem);

                _initUpload(newItem, dir, 1, function (res) {
                    $(newItem).find('td').find('img').attr('src', res.data.src);
                });

                _refreshAddStatus(box, count);
                _refreshSrc(inputObj, box);

                $(newItem).find('.del').click(function () {
                    preventBubble();
                    $(newItem).remove();
                    _refreshAddStatus(box, count);
                    _refreshSrc(inputObj, box);
                });

                $(newItem).find('.view').click(function () {
                    preventBubble();
                });
            });

            $(box).find('.item').each(function (r, row) {
                _initUpload($(row).find('td'), dir, 1, function (res) {
                    $(row).find('td').find('img').attr('src', res.data.path);
                    _refreshSrc(inputObj, box);
                    _refreshAddStatus(box, count);
                });

                $(row).find('.del').click(function () {
                    preventBubble();
                    $(row).remove();
                    _refreshAddStatus(box, count);
                    _refreshSrc(inputObj, box);
                });

                $(row).find('.view').click(function () {
                    preventBubble();
                });
            })


        })
    }
})(jQuery);

//上传文件
!(function ($) {

    function preventBubble(event) {
        var e = arguments.callee.caller.arguments[0] || event; //若省略此句，下面的e改为event，IE运行可以，但是其他浏览器就不兼容
        if (e && e.stopPropagation) {
            e.stopPropagation();
        } else if (window.event) {
            window.event.cancelBubble = true;
        }
    }

    var upload = layui.upload;

    function _boxTemp() {
        var temp = '<div class="upload-images-box"><ul></ul></div>';
        return $(temp);
    }

    function _itemTemp(src) {
        var temp = '<li class="item file">' +
            '               <div ><span class="file-src">' + src + '</span>' +
            '                   <a target="_blank" href="' + src + '" class="ico view fa fa-search-plus"></a><i class="ico del fa fa-trash"></i>' +
            '               </div>' +
            '           </li>';
        return $(temp);
    }

    function _addTemp() {
        var temp = '<li class="add">' +
            '           <div>' +
            '               <i class="fa fa-file-o"></i>' +
            '               <span>点击上传</span>' +
            '           </div>' +
            '       </li>';
        return $(temp);
    }

    function _initUpload(obj, dir, count, fn) {
        obj.on("click", function () {
            fileService(function (files) {

                files.map(function (f, i) {
                    var d = {data: {path: f.path, src: f.src}};
                    fn(d);
                });

            }, count > 1 ? true : false);
        });
    }

    function _refreshAddStatus(box, count) {
        if (box.find('.item').length >= count)
            $(box).find('.add').hide();
        else
            $(box).find('.add').show();
    }

    function _refreshSrc(obj, box) {
        var src = $(box).find('.file-src').html();
        $(obj).val(src);
    }

    $.fn.uploadFile = function (flag) {
        flag = flag || false;
        $(this).each(function (i, inputObj) {

            var count = (typeof $(inputObj).attr("count") == "undefined") ? 1 : $(inputObj).attr("count");
            var dir = $(inputObj).attr("dir");

            var box = _boxTemp();

            var srcs = $(inputObj).val();
            var arrSrc = srcs.split(',');
            $(arrSrc).each(function (k, src) {
                if ($.trim(src) != '')
                    box.find('ul').append(_itemTemp(src));
            });

            box.find('ul').append(_addTemp());
            _refreshAddStatus(box, count);

            $(inputObj).parent().append(box);


            _initUpload($(box).find('.add > div'), dir, count, function (res) {
                var newItem = _itemTemp(flag ? res.data.src : res.data.path);

                $(box).find('.add').before(newItem);

                _initUpload(newItem, dir, 1, function (res) {
                    $(newItem).find('td').find('img').attr('src', res.data.src);
                });

                _refreshAddStatus(box, count);
                _refreshSrc(inputObj, box);

                $(newItem).find('.del').click(function () {
                    preventBubble();
                    $(newItem).remove();
                    _refreshAddStatus(box, count);
                    _refreshSrc(inputObj, box);
                });

                $(newItem).find('.view').click(function () {
                    preventBubble();
                });
            });

            $(box).find('.item').each(function (r, row) {
                _initUpload($(row).find('td'), dir, 1, function (res) {
                    _refreshSrc(inputObj, box);
                    _refreshAddStatus(box, count);
                });

                $(row).find('.del').click(function () {
                    preventBubble();
                    $(row).remove();
                    _refreshAddStatus(box, count);
                    _refreshSrc(inputObj, box);
                });

                $(row).find('.view').click(function () {
                    preventBubble();
                });
            })


        })
    }
})(jQuery);