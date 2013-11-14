'use strict'

// App Module

var app = angular.module('martianblog', ['ui.router']);
app.config(['$stateProvider', function ($stateProvider) {
	$stateProvider.
		state('front', {
			url: "",
			templateUrl: "static/template/front/template.html",
			controller: "FrontCtrl"
		}).
		state('front.list', {
			url: "/articles",
			templateUrl: "static/template/front/articleList.html",
			controller: "ArticleListCtrl"
		}).
		state('front.detail', {
			url: "/articles/:articleId",
			templateUrl: "static/template/front/articleDetail.html",
			controller: "ArticleDetailCtrl"
		});
}]);