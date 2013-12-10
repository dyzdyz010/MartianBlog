'use strict'

// Controllers

function FrontCtrl ($scope, $http, $state) {
	$scope.title = 'Moonlightter';
	$scope.subtitle = "This is subtitle.";
}

function ArticleListCtrl ($scope, $http) {
	$http.get('/articles', {params: {status: "published"}}).success(function (data) {
		console.log(data);
		for (var i in data) {
			data[i].date = moment(data[i].date, 'YYYY-MM-DD');
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

function AdminCtrl ($scope, $http) {
	
}

function DashboardCtrl ($scope, $http) {
	
}

function AdminArticleListCtrl ($scope, $http) {
	$http.get('/articles').success(function (data) {
		$scope.articles = data;
	});
}

function AdminArticleEditCtrl ($scope, $http, $stateParams, $state, $notification) {
	var postUrl = '/admin/post';
	marked.setOptions({
		highlight: function (code, lang) {
			console.log('lang: '+lang+', code: '+code);
			return (lang == undefined) ? hljs.highlightAuto(code).value : hljs.highlight(lang, code).value;
		}
	});

	if ($stateParams.articleId != undefined) {
		console.log($stateParams.articleId);
		$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
			$scope.article = data
			$scope.date = moment(data.date).format('YYYY-MM-DD');
		});

		postUrl = postUrl + '/update';
	} else {
		postUrl = postUrl + '/new';
	}

	// Save draft form button action
	$scope.draft = function () {
		var article = $scope.article;
		console.log(article);
		article.status = 'draft';
		article.date = generateDate($scope.date).format();
		if (article.date != undefined) {
			article.author = 'DYZ';
			console.log(article);
			$http.post('/admin/post/new', article).success(function (data, status) {
				console.log(data);
				if (data.code == 200) {
					$state.transitionTo('admin.articles');
					$notification.success('Success', 'Save draft succeed.');
				} else{
					$notification.error('Error', 'Something is wrong.');
				}
			})
		} else {

		}
	};

	// Post form button action
	$scope.post = function () {
		article.status = 'published';
	}

	var generateDate = function (d) {
		var date = moment();
		if (d == undefined || d == '') {
			return date;
		}

		var temp = moment(d, 'YYYY-MM-DD', true);
		if (!temp.isValid()) {
			return undefined;
		}

		date.date(temp.date());
		date.month(temp.month());
		date.year(temp.year());

		return date;
	}
}