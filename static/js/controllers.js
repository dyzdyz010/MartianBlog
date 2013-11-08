'use strict'

// Controllers

function ArticlesListCtrl ($scope, $http) {
	$http.get('/articles').success(function (data) {
		$scope.articles = data;
	});
}