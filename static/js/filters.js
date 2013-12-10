'use strict';

// Filters

var filters = angular.module('martianblogFilters', ['ngSanitize']);

filters.filter('timeago', function() {
	return function (input) {
		return moment(input, 'YYYY-MM-DD').fromNow();
	};
});

filters.filter('markdown', function ($sce) {
	return function (input) {
		return $sce.trustAsHtml(marked((input==undefined)?"":input));
	};
});