'use strict'

// Controllers

function FrontCtrl ($scope, $http, $state) {
	$scope.template = 'static/template/front/template.html';
	$scope.title = 'Moonlightter';
	$scope.subtitle = "This is subtitle.";

	if ($state.current.name == 'front') {
		$state.transitionTo('front.list');
	} else {
		$state.transitionTo($state.current);
	};
	console.log($state.current);
}

function ArticleListCtrl ($scope, $http) {
	$http.get('/articles').success(function (data) {
		$scope.articles = data;
	});
}

function ArticleDetailCtrl ($scope, $http, $stateParams) {
	$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
		$scope.article = data
	})
}

function AdminCtrl ($scope, $http) {
	
}