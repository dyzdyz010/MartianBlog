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
	$http.get('/articles', {params: {status: "published"}}).success(function (data) {
		console.log(data);
		for (var i in data) {
			data[i].Date = moment(data[i].Date, 'YYYY-MM-DD');
		}
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
	var postUrl = '/admin/post';

	marked.setOptions({
		highlight: function (code, lang) {
			console.log('lang: '+lang+', code: '+code);
			return (lang == undefined) ? hljs.highlightAuto(code).value : hljs.highlight(lang, code).value;
		}
	});
	console.log('AdminArticleEditCtrl');
	if ($stateParams.articleId != undefined) {
		console.log($stateParams.articleId);
		$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
			$scope.article = data
		});

		postUrl = postUrl + '/update';
	} else {
		postUrl = postUrl + '/new';
	}

	$scope.draft = function (article) {
		article.Status = 'draft';
		article.Date = moment().format('YYYY-MM-DD');
		if (true) {
			article.Id = 1000;
			article.Author = 'DYZ';
			console.log(article);
			$http.post('/admin/post/new', article).success(function (data, status) {
				console.log(status);
			})
		};
	};

	$scope.post = function (article) {
		article.Status = 'published';
	}

	function dateValid (date) {
		return date != '' && moment(date, 'YYYY-MM-DD').isValid() || date == '';
	}

	console.log(moment('YYYY-MM-DD'));
}