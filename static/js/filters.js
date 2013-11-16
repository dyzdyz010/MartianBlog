'use strict';

// Filters

var filters = angular.module('martianblogFilters', ['ngSanitize']);

filters.filter('calendar', function() {
	return function (input) {
		console.log(input);
		return moment(input, 'YYYY-MM-DD').calendar();
	};
});

filters.filter('markdown', function ($sce) {
	return function (input) {
		return $sce.trustAsHtml(marked(input));
	};
});