$(function () {
    var clock = $('#clock'), alarm = clock.find('.alarm'), ampm = clock.find('.ampm'),
        dialog = $('#alarm-dialog').parent(), alarm_set = $('#alarm-set'), alarm_clear = $('#alarm-clear'),
        time_is_up = $('#time-is-up').parent();
    var alarm_counter = -1;
    var digit_to_name = 'zero one two three four five six seven eight nine'.split(' ');
    var digits = {};
    var positions = ['h1', 'h2', ':', 'm1', 'm2', ':', 's1', 's2'];
    var digit_holder = clock.find('.digits');
    $.each(positions, function () {
        if (this == ':') {
            digit_holder.append('<div class="dots">');
        }
        else {
            var pos = $('<div>');
            for (var i = 1; i < 8; i++) {
                pos.append('<span class="d' + i + '">');
            }
            digits[this] = pos;
            digit_holder.append(pos);
        }
    });
    var weekday_names = 'MON TUE WED THU FRI SAT SUN'.split(' '), weekday_holder = clock.find('.weekdays');
    $.each(weekday_names, function () {
        weekday_holder.append('<span>' + this + '</span>');
    });
    var weekdays = clock.find('.weekdays span');
    (function update_time() {
        var now = moment().format("hhmmssdA");
        digits.h1.attr('class', digit_to_name[now[0]]);
        digits.h2.attr('class', digit_to_name[now[1]]);
        digits.m1.attr('class', digit_to_name[now[2]]);
        digits.m2.attr('class', digit_to_name[now[3]]);
        digits.s1.attr('class', digit_to_name[now[4]]);
        digits.s2.attr('class', digit_to_name[now[5]]);
        var dow = now[6];
        dow--;
        if (dow < 0) {
            dow = 6;
        }
        weekdays.removeClass('active').eq(dow).addClass('active');
        ampm.text(now[7] + now[8]);
        if (alarm_counter > 0) {
            alarm_counter--;
            alarm.addClass('active');
        }
        else if (alarm_counter == 0) {
            time_is_up.fadeIn();
            try {
                $('#alarm-ring')[0].play();
            }
            catch (e) {
            }
            alarm_counter--;
            alarm.removeClass('active');
        }
        else {
            alarm.removeClass('active');
        }
        setTimeout(update_time, 1000);
    })();
    $('#switch-theme').click(function () {
        clock.toggleClass('light dark');
    });
    $('.alarm-button').click(function () {
        dialog.trigger('show');
    });
    dialog.find('.close').click(function () {
        dialog.trigger('hide')
    });
    dialog.click(function (e) {
        if ($(e.target).is('.overlay')) {
            dialog.trigger('hide');
        }
    });
    alarm_set.click(function () {
        var valid = true, after = 0, to_seconds = [3600, 60, 1];
        dialog.find('input').each(function (i) {
            if (this.validity && !this.validity.valid) {
                valid = false;
                this.focus();
                return false;
            }
            after += to_seconds[i] * parseInt(parseInt(this.value));
        });
        if (!valid) {
            alert('Please enter a valid number!');
            return;
        }
        if (after < 1) {
            alert('Please choose a time in the future!');
            return;
        }
        alarm_counter = after;
        dialog.trigger('hide');
    });
    alarm_clear.click(function () {
        alarm_counter = -1;
        dialog.trigger('hide');
    });
    dialog.on('hide', function () {
        dialog.fadeOut();
    }).on('show', function () {
        var hours = 0, minutes = 0, seconds = 0, tmp = 0;
        if (alarm_counter > 0) {
            tmp = alarm_counter;
            hours = Math.floor(tmp / 3600);
            tmp = tmp % 3600;
            minutes = Math.floor(tmp / 60);
            tmp = tmp % 60;
            seconds = tmp;
        }
        dialog.find('input').eq(0).val(hours).end().eq(1).val(minutes).end().eq(2).val(seconds);
        dialog.fadeIn();
    });
    time_is_up.click(function () {
        time_is_up.fadeOut();
    });
});