'use strict'

// Controllers

function FrontCtrl ($scope, $http, $state, $rootScope) {
	$scope.title = 'Moonlightter';
	$scope.subtitle = "This is subtitle.";

	// if ($state.current.name == 'front') {
	// 	console.log('This is front state, prepare to front.list state.');
	// 	$state.transitionTo('front.list');
	// }
}

function ArticleListCtrl ($scope, $http) {
	$http.get('/articles').success(function (data) {
		$scope.articles = data;
	});
}

function ArticleDetailCtrl ($scope, $http, $stateParams) {
	console.log('ArticleDetailCtrl');
	$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
		$scope.article = data;
	});
}

function AdminCtrl ($scope, $http, $state) {
	if ($state.current.name == 'admin') {
		console.log('This is admin state, prepare to admin.dashboard state.');
		$state.transitionTo('admin.dashboard');
	}
}

function DashboardCtrl ($scope, $http) {
	
}

function AdminArticleListCtrl ($scope, $http) {
	$http.get('/articles').success(function (data) {
		$scope.articles = data;
	});
}

function AdminArticleEditCtrl ($scope, $http, $stateParams) {
	console.log('ArticleDetailCtrl');
	if ($stateParams.articleId != undefined) {
		$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
			$scope.article = data
		});
	};
}