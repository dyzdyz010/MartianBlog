'use strict'

/* Filters */

angular.module('markdownFilters', []).filter('markdown', function() {
	return function(input) {
		if (input == undefined) {
			return "";
		};
		var converter = new Showdown.converter();
		var htmlStr = converter.makeHtml(input);
		return htmlStr;
	};
});