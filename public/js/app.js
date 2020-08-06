// Initialize tooltips
$(document).ready(function() {
    $('[data-toggle="tooltip"]').tooltip();
});

// Initialize confirm alerts
$(document).ready(function() {
    $('a[data-confirm="confirm"]').on('click', function(event) {
        var $this = $(this);
        event.preventDefault();
        bootbox.confirm({
            message: $(this).attr('data-message'),
            buttons: {
                confirm: {
                    label: 'Yes'
                },
                cancel: {
                    label: 'No'
                }
            },
            callback: function(result) {
                if (result) {
                    window.location = $this.attr('href');
                }
            }
        });
    });
});
