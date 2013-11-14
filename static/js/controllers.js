'use strict'

// Controllers

function FrontCtrl ($scope, $http, $state) {
	console.log('FrontCtrl: ' + $state.current);
	$scope.template = 'static/template/front/template.html';
	$scope.title = 'Moonlightter';
	$scope.subtitle = "This is subtitle.";

	if ($state.current.name == 'front') {
		console.log('This is front state, prepare to front.list state.');
		$state.transitionTo('front.list');
	}
}

function ArticleListCtrl ($scope, $http) {
	$http.get('/articles').success(function (data) {
		$scope.articles = data;
	});
}

function ArticleDetailCtrl ($scope, $http, $stateParams) {
	console.log('ArticleDetailCtrl');
	$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
		$scope.article = data
	})
}

function AdminCtrl ($scope, $http) {
	
}