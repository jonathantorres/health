
/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

require('./bootstrap');

window.Vue = require('vue');
window.bootbox = require('bootbox');

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

/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

Vue.component('example', require('./components/Example.vue'));

const app = new Vue({
    el: '#app'
});
