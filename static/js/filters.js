'use strict';

// Filters

angular.module('martianblogFilters', []).filter('calendar', function() {
  return function(input) {
    return moment(input, 'YYYY-MM-DD').calendar();
  };
});
