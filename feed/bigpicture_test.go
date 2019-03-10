package feed

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

const file = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<script type="text/javascript">
	var methode = {};
	window.globe = window.globe || {};
	globe.version = '19745I2111';
</script>
<!-- /news/bigpicture auto:  -->
	<title>Globe staff photos of the month, February 2019 - The Boston Globe</title>
<meta name="eomportal-instanceid" content="217"/>
<meta name="eomportal-id" content="1160347"/>
<meta name="eomportal-loid" content="5.0.1114829766"/>
<meta name="eomportal-uuid" content="a91542fe-2699-11e9-91c7-fd316211b708"/>
<meta name="eomportal-lastUpdate" content="Sat Mar 09 19:51:19 EST 2019"/>
<title>Globe staff photos of the month, February 2019 - The Boston Globe</title>
<!-- Methode uuid: "a91542fe-2699-11e9-91c7-fd316211b708" -->
<script src="//cdn.blueconic.net/bostonglobemedia.js"></script>

	<meta property="og:type" content="article"/>
	<meta property="og:description" content="Here’s a look at some of the best images taken by Globe photographers last month: winter weather, presidential candidates on the stump, a Chinese New Year celebration, Red Sox spring training, and the Patriots’ Super Bowl LIII victory parade. "/>
	<meta property="og:title" content="Globe staff photos of the month, February 2019 - The Boston Globe"/>
	<meta property="og:site_name" content="BostonGlobe.com" />
	<meta property="fb:admins" content="100001044693856" />
	<meta property="fb:admins" content="507486035" />
	<meta property="fb:app_id" content="103933749691726" />
	<meta property="fb:pages" content="5637143257" />
	<meta property="article:publisher" content="https://www.facebook.com/globe" />
	<link rel='apple-touch-icon' sizes='57x57' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-57x57.png'>
	<link rel='apple-touch-icon' sizes='60x60' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-60x60.png'>
	<link rel='apple-touch-icon' sizes='72x72' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-72x72.png'>
	<link rel='apple-touch-icon' sizes='76x76' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-76x76.png'>
	<link rel='apple-touch-icon' sizes='114x114' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-114x114.png'>
	<link rel='apple-touch-icon' sizes='120x120' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-120x120.png'>
	<link rel='apple-touch-icon' sizes='144x144' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-144x144.png'>
	<link rel='apple-touch-icon' sizes='152x152' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-152x152.png'>
	<link rel='apple-touch-icon' sizes='180x180' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/apple-touch-icon-180x180.png'>

	<link rel='icon' type='image/png' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/favicon-16x16.png' sizes='16x16'>
	<link rel='icon' type='image/png' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/favicon-32x32.png' sizes='32x32'>
	<link rel='icon' type='image/png' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/favicon-96x96.png' sizes='96x96'>
	<link rel='icon' type='image/png' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/favicon-194x194.png' sizes='194x194'>
	<link rel='icon' type='image/png' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/favicon-192x192.png' sizes='192x192'>

	<link rel='mask-icon' href='/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/safari-pinned-tab.svg' color='#000000'>

	<link rel="shortcut icon" type="image/gif" href="/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/favicon.ico" />

	<meta name='msapplication-TileColor' content='#ffffff'>
	<meta name='msapplication-TileImage' content='//c.o0bg.com/rw/SysConfig/WebPortal/BostonGlobe/Framework/bg-images/png/favicon/mstile-144x144.png'>
	<meta name='theme-color' content='#ffffff'>

	<!--[if gte IE 9 ]>
	<meta name="application-name" content="Homepage"/>
	<meta name="msapplication-tooltip" content="This is the tooltip"/>
	<meta name="msapplication-starturl" content="/?comp=9"/>
	<![endif]-->

	<!--[if lte IE 8 ]>
	<meta name="application-name" content="Boston Globe"/>
	<meta name="msapplication-TileColor" content="#000000"/>
	<meta name="msapplication-TileImage" content="25a75e8d-4542-40a4-80aa-253aea7efc28.png"/>
	<link rel="stylesheet" media="print" href="/css/globe-ie-print.css?v=19745I2111" type="text/css" />
	<![endif]-->

	<meta property="twitter:account_id" content="95431448" />
<meta name="robots" content="noarchive" />
<meta name="description" content="Here’s a look at some of the best images taken by Globe photographers last month: winter weather, presidential candidates on the stump, a Chinese New Year celebration, Red Sox spring training, and the Patriots’ Super Bowl LIII victory parade. " />
<meta name="keywords" content="" />
<meta name="news_keywords" content="" />
<meta property="og:url" content="https://www.bostonglobe.com/news/bigpicture/2019/03/02/globe-staff-photos-month-february/zsbKlFlxkRDAiLYNucIwxJ/story.html"/>
<link rel="canonical" href="https://www.bostonglobe.com/news/bigpicture/2019/03/02/globe-staff-photos-month-february/zsbKlFlxkRDAiLYNucIwxJ/story.html" />
<meta name="twitter:card" content="summary_large_image" />
<meta name="twitter:site" content="BostonGlobe" />
<meta name="twitter:description" content="Here’s a look at some of the best images taken by Globe photographers last month: winter weather, presidential candidates on the stump, a Chinese New Year celebration, Red Sox spring training, and the Patriots’ Super Bowl LIII victory parade. "/>
	<meta name="twitter:title" content="Globe staff photos of the month, February 2019 - The Boston Globe"/>
	<meta name="twitter:image" content="https://www.bostonglobe.com/rw/SysConfig/WebPortal/BostonGlobe/Framework/images/logo-bg-twitter.jpg" />
	<meta property="og:image" content="//c.o0bg.com/rw/SysConfig/WebPortal/BostonGlobe/Framework/images/logo-bg.jpg" />

<meta property="master-uuid" content="a91542fe-2699-11e9-91c7-fd316211b708"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
<meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" /> <!--[if gte IE 6]><link href="/css/html5reset.css,globe-globals.css,globe-masthead.css,globe-nav.css,globe-nav-menus.css,globe-main.css,globe-footer.css,globe-print.css" rel="stylesheet"><![endif]-->
<!--[if !IE]><link href="/css/html5reset.css,globe-globals.css,globe-masthead.css,globe-nav.css,globe-nav-menus.css,globe-main.css,globe-footer.css,globe-print.css?v=19745I2111" rel="stylesheet" media="only all"><![endif]-->
<script src="/js/lib/rwd-images.js,lib/respond.min.js,lib/modernizr.custom.min.js,globe-define.js,globe-controller.js?v=19745I2111"></script>
	<script type="text/javascript" title="https://iconizr.com | © 2014 Joschi Kuphal | MIT">!function(){var a=window,b=a.document,c=arguments,d="createElementNS",e="http://www.w3.org/",f="svg",g=function(a){var g=b.getElementsByTagName("script")[0],h=b.createElement("link");h.rel="stylesheet",h.type="text/css",h.href=c[1*a|2*(b[d]&&b[d](e+"2000/"+f,f).createSVGRect&&b.implementation.hasFeature(e+"TR/SVG11/feature#Image","1.1"))],g.parentNode.insertBefore(h,g)},h=new a.Image;h.onerror=function(){g(0)},h.onload=function(){g(1===h.width&&1===h.height)},h.src="data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///ywAAAAAAQABAAACAUwAOw=="}("/css/icon-png-sprite.css","/css/icon-png-data.css","/css/icon-svg-sprite.css","/css/icon-svg-data.css");</script>
	<noscript>
	   <link href="/css/icon-png-sprite.css?v=19745I2111" rel="stylesheet" type="text/css" media="all">
	</noscript>

<meta name="datePublished" content="2019-03-02T14:33:56.000Z" />

<!-- /news/bigpicture and big-picture -->
<script src="/js/lib/jquery.js,lib/lo-dash-custom-2.4.1.js,lib/a9.js,lib/pb.js,dist/ad-init.js,globe-newsletter.js,globe-profile-page.js,dist/globe-topic-nav.js,dist/rakuten.js?v=19745I2111"></script>
<link rel="stylesheet" type="text/css" media="all" href="/css/html5reset.css,globe-globals.css,dist/site-header.css,globe-main.css,globe-footer.css,globe-polldaddy.css,article.css,gallery.css,globe-topic-nav.css,globe-profile-page.css,globe-bgbrandlab-belt.css?v=19745I2111"/>
<link rel="stylesheet" type="text/css" media="print" href="/css/globe-print.css?v=19745I2111"/>
<!-- oas_sitepage:/news/bigpicture/homepage -->
<!-- webtype is big-picture section is  adPageType  oas_sitepage /news/bigpicture/homepage-->

<!--  lowestRank:  taxonomy:  section:  -->

		<!-- admode:  smode user -->
		<!-- modAdunit: bostonglobe.com/news/bigpicture  -->
    <script type="text/javascript" src="//dz9qn8fh4jznm.cloudfront.net/script.js"></script> 

    <script async class="kxint" data-namespace="bostonglobe" type="text/javascript">
        window.Krux||((Krux=function(){Krux.q.push(arguments)}).q=[]),function(){function n(n){var r="kxbostonglobe_"+n,u=function(){try{return window.localStorage}catch(n){return null}}();if(u)return u[r]||"";if(navigator.cookieEnabled){var e=document.cookie.match(r+"=([^;]*)");return e&&unescape(e[1])||""}return""}Krux.user=n("user"),Krux.segments=n("segs")?n("segs").split(","):[]}();
    </script>

    <script type="text/javascript">
        var gads = document.createElement("script");
        gads.async = true;
        gads.type = "text/javascript";
        gads.src = "//www.googletagservices.com/tag/js/gpt.js";
        var node = document.getElementsByTagName("script")[0];
        node.parentNode.insertBefore(gads, node);
    </script>

    <script>
        var googletag = googletag || {};
        googletag.cmd = googletag.cmd || [];

        googletag.cmd.push(function() {
            googletag.pubads().disableInitialLoad();
        });
    </script>

    <script>
    	  var globe = globe || {};
	      globe.dfp = {
	        networkCode: 61381659
	      };

          globe.dfp.keyValuePairs = {
            s1: 'bostonglobe.com',
            s2: 'news',
            s3: 'bigpicture',
            s4: '',
            pgtype: '',
            pageurl: '_news_bigpicture_2019_03_02_globe-staff-photos-month-february_zsbKlFlxkRDAiLYNucIwxJ_story_html',
            meta: "",
            weather: "",
            temp: "",
            aam: ""
        }
        globe.dfp.adUnit = 'bostonglobe.com/news/bigpicture';
        globe.dfp.isBigArticle = ( 'Gallery' === 'Special Article' ? true : false );
    </script>
<script class="kxct" data-id="skwb6bm8t" data-timing="async" data-version="3.0" type="text/javascript">
window.Krux||((Krux=function(){Krux.q.push(arguments)}).q=[]);
(function(){
   var k=document.createElement('script');k.type='text/javascript';k.async=true;
   k.src=(location.protocol==='https:'?'https:':'http:')+'//cdn.krxd.net/controltag/skwb6bm8t.js';
   var s=document.getElementsByTagName('script')[0];s.parentNode.insertBefore(k,s);
}());
</script>
<script src="//dc8xl0ndzn2cb.cloudfront.net/js/bostonglobe/v0/keywee.min.js" type="text/javascript" async></script>
</head>

<body itemscope itemtype="http://schema.org/WebPage" class='type-internal big-picture'>

<div id="contain">

<div id="contain">
	<!--
	type: EOM::MediaGallery
	section: /news/bigpicture
	topSection: /news/bigpicture
	autoSection: 
	sectionTitle: The Big Picture
-->
<header class="site-header site-header--section js-site-header">

	<aside class="takeover-menu js-takeover-menu">
		<div class="takeover-menu__wrap js-takeover-menu-wrap">

			<div class="js-user-settings" data-location="Menu" data-baseurl="https://www.bostonglobe.com"></div>
		<div class="js-user-subscribe" data-location="Menu" data-baseurl="https://www.bostonglobe.com"></div>
	<div class="takeover-menu__content js-takeover-menu-content">
				<nav class="takeover-menu__section takeover-menu__section--left">
					<ul class="takeover-menu__list" role="menu">
						<li class="takeover-menu__item">
										<a href="/metro?p1=BGMenu" class="takeover-menu__link">Metro</a>
									</li>
								<li class="takeover-menu__item">
										<a href="/sports?p1=BGMenu" class="takeover-menu__link">Sports</a>
									</li>
								<li class="takeover-menu__item">
										<a href="/business?p1=BGMenu" class="takeover-menu__link">Business & Tech</a>
									</li>
								<li class="takeover-menu__item">
										<a href="/opinion?p1=BGMenu" class="takeover-menu__link">Opinion</a>
									</li>
								<li class="takeover-menu__item">
										<a href="/news/politics?p1=BGMenu" class="takeover-menu__link">Politics</a>
									</li>
								<li class="takeover-menu__item">
										<a href="/lifestyle?p1=BGMenu" class="takeover-menu__link">Lifestyle</a>
									</li>
								<li class="takeover-menu__item">
										<a href="/news/marijuana?p1=BGMenu" class="takeover-menu__link">Marijuana</a>
									</li>
								<li class="takeover-menu__item">
										<a href="/arts?p1=BGMenu" class="takeover-menu__link">Arts</a>
									</li>
								<li class="takeover-menu__item">
										<a href="http://www.boston.com/section/cars?s_campaign=bg:hp:mainnav:cars" class="takeover-menu__link"
										target="_blank"
										onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Menu - Nav - Cars');">
										Cars</a>
									</li>
								<li class="takeover-menu__item">
										<a href="http://realestate.boston.com?s_campaign=bg:hp:mainnav:realestate" class="takeover-menu__link"
										target="_blank"
										onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Menu - Nav - Real Estate');">
										Real Estate</a>
									</li>
								<li class="takeover-menu__item">
										<a href="https://globeevents.splashthat.com?s_campaign=bg:hp:mainnav:splashthat" class="takeover-menu__link"
										target="_blank"
										onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Menu - Nav - Events');">
										Events</a>
									</li>
								</ul>
				</nav>

				<div class="takeover-menu__section takeover-menu__section--right">
					<div class="takeover-menu__suggestion-header">Most popular on BostonGlobe.com</div>

	<ul class="takeover-menu__suggestion-list">
		<li class="takeover-menu__suggestion-item">
				<a href="//www.bostonglobe.com/metro/2019/03/09/drugs-took-their-children-but-not-their-hope-that-others-might-saved/42PBgLkAx3q5Db9Khno8wM/story.html?p1=BGMenu_Article" class="takeover-menu__link" aria-label="recommended article">
							<div class="takeover-menu__suggestion-img-wrap">
								<img class="takeover-menu__suggestion-img" src="//bostonglobe.com/rf/image_585w/Boston/2011-2020/2019/03/05/BostonGlobe.com/Metro/Images/greenhouseParents2.jpg"/>
							</div>
							<div class="takeover-menu__suggestion-headline">Road to Recovery: Drugs took their children, but not their hope that others might be saved</div>
		 				</a>
					</li>
		<li class="takeover-menu__suggestion-item">
				<a href="//www.bostonglobe.com/metro/2019/03/04/the-backyard-mechanic-who-taking-tesla/Sv1l8q2sxpQvTFMp13VFwM/story.html?p1=BGMenu_Article" class="takeover-menu__link" aria-label="recommended article">
							<div class="takeover-menu__suggestion-img-wrap">
								<img class="takeover-menu__suggestion-img" src="//bostonglobe.com/rf/image_585w/Boston/2011-2020/2019/03/03/BostonGlobe.com/Metro/Images/wiggs_Tesla_04.jpg"/>
							</div>
							<div class="takeover-menu__suggestion-headline">The backyard mechanic who is taking on Tesla</div>
		 				</a>
					</li>
		<li class="takeover-menu__suggestion-item">
				<a href="//www.bostonglobe.com/metro/2019/03/09/mass-authorities-face-steep-hurdles-shutting-down-sex-trafficking/yFW40bBEEN9QtuIHaqTrmM/story.html?p1=BGMenu_Article" class="takeover-menu__link" aria-label="recommended article">
							<div class="takeover-menu__suggestion-img-wrap">
								<img class="takeover-menu__suggestion-img" src="//bostonglobe.com/rf/image_585w/Boston/2011-2020/2019/03/06/BostonGlobe.com/Metro/Images/KREITER03062019sextrafficarlington3.jpg"/>
							</div>
							<div class="takeover-menu__suggestion-headline">Mass. authorities face steep hurdles in shutting down sex trafficking</div>
		 				</a>
					</li>
		<li class="takeover-menu__suggestion-item">
				<a href="//www.bostonglobe.com/sports/patriots/2019/03/09/patriots-could-diving-into-quarterback-market/WrQo4lp16aQ4WqulGp8yNO/story.html?p1=BGMenu_Article" class="takeover-menu__link" aria-label="recommended article">
							<div class="takeover-menu__suggestion-img-wrap">
								<img class="takeover-menu__suggestion-img" src="//bostonglobe.com/rf/image_585w/Boston/2011-2020/2019/03/08/BostonGlobe.com/Sports/Images/38ea6505049b4aee81e2d7d1a57035d6-38ea6505049b4aee81e2d7d1a57035d6-0.jpg"/>
							</div>
							<div class="takeover-menu__suggestion-headline">Patriots could be diving into the quarterback market</div>
		 				</a>
					</li>
		</ul>
<form class="takeover-menu__search" action="/queryResult/search" method="get" role="search">
						<input class="takeover-menu__search-input" type="search" placeholder="Search BostonGlobe.com" name="q">
						<span class="takeover-menu__search-icon">
							<svg width="21" height="22" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="globe-search-title">
	<title id="globe-search-title">Search The Boston Globe</title>
	<g transform="translate(2.064 1.892)" stroke="#777" stroke-width="3" fill="none" fill-rule="evenodd"><ellipse cx="10.524" cy="6.545" rx="6.476" ry="6.545"/>
		<path d="M6.07 11.864L.406 17.59" stroke-linecap="square"/>
	</g>

	<!-- fallback .png -->
	<image src="/bg-images/png/search-magnifying-glass.png" xlink:href=""/>
</svg>
</span>
					</form>
				</div>

				<nav class="takeover-menu__section takeover-menu__section--footer">
					<ul class="takeover-menu__list" role="menu">
						<li class="takeover-menu__item takeover-menu__item--subsection">
								<a href="/todayspaper?p1=BGMenu_Subnav" class="takeover-menu__link takeover-menu__link--subsection">Today's Paper</a>
							</li>
						<li class="takeover-menu__item takeover-menu__item--subsection">
								<a href="/magazine?p1=BGMenu_Subnav" class="takeover-menu__link takeover-menu__link--subsection">Magazine</a>
							</li>
						<li class="takeover-menu__item takeover-menu__item--subsection">
								<a href="/metro/obituaries?p1=BGMenu_Subnav" class="takeover-menu__link takeover-menu__link--subsection">Obituaries</a>
							</li>
						<li class="takeover-menu__item takeover-menu__item--subsection">
								<a href="http://www.boston.com/weather?p1=BGMenu_Subnav" class="takeover-menu__link takeover-menu__link--subsection">Weather</a>
							</li>
						<li class="takeover-menu__item takeover-menu__item--subsection">
								<a href="/lifestyle/comics?p1=BGMenu_Subnav" class="takeover-menu__link takeover-menu__link--subsection">Comics</a>
							</li>
						<li class="takeover-menu__item takeover-menu__item--subsection">
								<a href="/lifestyle/crossword?p1=BGMenu_Subnav" class="takeover-menu__link takeover-menu__link--subsection">Crossword</a>
							</li>
						<li class="takeover-menu__item takeover-menu__item--subsection">
								<a href="/news/bigpicture?p1=BGMenu_Subnav" class="takeover-menu__link takeover-menu__link--subsection">The Big Picture</a>
							</li>
						</ul>

					<form class="takeover-menu__search takeover-menu__search--footer" action="/queryResult/search" method="get" role="search">
						<input class="takeover-menu__search-input" type="search" placeholder="Search BostonGlobe.com" name="q">
						<span class="takeover-menu__search-icon">
							<svg width="21" height="22" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="globe-search-title">
	<title id="globe-search-title">Search The Boston Globe</title>
	<g transform="translate(2.064 1.892)" stroke="#777" stroke-width="3" fill="none" fill-rule="evenodd"><ellipse cx="10.524" cy="6.545" rx="6.476" ry="6.545"/>
		<path d="M6.07 11.864L.406 17.59" stroke-linecap="square"/>
	</g>

	<!-- fallback .png -->
	<image src="/bg-images/png/search-magnifying-glass.png" xlink:href=""/>
</svg>
</span>
					</form>
				</nav>
			</div>
		</div>
	</aside>

	<div class="site-header__menu js-header-menu">
		<a class="site-header__hamburger js-hamburger-menu" href="#" role="button" aria-label="main menu" onclick="var s=s_gi('nytbostonglobecom');s.tl(this,'o','BG Menu - Open');">
			<span class="site-header__hamburger-line takeover-menu__hamburger-line"></span>
			<span class="site-header__hamburger-text">Menu</span>
		</a>
		<a href="/?p1=BGHeader_Logo" class="site-header__home-link takeover-menu__home-link" aria-label="Visit The Boston Globe homepage">
			<span class="site-header__logo-link">
				<svg class="bg-logo bg-logo--full" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="globe-logo--big" x="0px" y="0px" width="140px" height="17px" viewBox="0 0 140 17" enable-background="new 0 0 140 18.674" xml:space="preserve" aria-labelledby="globe-title">
	<title id="globe-title">The Boston Globe</title>
	<path d="M129.334,11.479c0,1.375-0.062,4.065-0.062,4.16l-0.095,0.094l-2.597-1.625V7.727c0.406-0.625,0.97-0.938,1.471-0.938   c0.594,0,1.03,0.438,1.249,1.032C129.3,8.071,129.334,9.729,129.334,11.479 M137.902,8.885l-2.408,1.97V6.383L137.902,8.885z    M131.521,14.546c-0.156-1.406-0.123-5.225-0.123-7.006c0-1.157-1.004-2.002-2.254-2.534l-2.564,2.158V2.472   c0-0.781,0.47-1.626,1.031-1.626V0.689c-1.219,0-2.563,0.971-2.783,1.597c-0.28-0.47-0.938-0.752-1.562-0.752l-0.031,0.25   c0.625,0,1.221,1.064,1.221,2.941c0,3.784-0.033,7.413-0.033,8.974c0,0.503-0.062,0.847-0.344,1.253l3.849,2.345   C129.395,15.923,129.958,15.517,131.521,14.546 M140,13.952l-0.189-0.189c-0.97,0.971-1.533,1.377-2.095,1.377   c-0.563,0-1.096-0.345-2.222-1.22v-2.563l4.41-3.536c-1.032-0.875-3.158-3.002-3.158-3.002c-1.189,1.064-1.878,1.658-3.535,2.815   c0.062,0.344,0.125,0.812,0.125,1.188v4.973c0,0.345-0.031,0.595-0.156,1l3.473,2.438L140,13.952z M98.868,11.979   c0,0.846-0.252,1.377-0.813,1.814c-1.281-1.251-1.906-3.003-1.906-4.754c0-2.626,0.906-3.629,2.158-4.535   c0.188-0.157,0.375-0.282,0.562-0.406V11.979z M102.247,15.576c-1.658-0.124-2.938-0.687-3.94-1.531   c1.529-1.062,2.502-1.594,2.502-3.223v-8.1l0.469-0.374c0.345,0.124,0.625,0.249,0.971,0.374L102.247,15.576L102.247,15.576z    M107.125,11.886c0,0.909-0.094,1.534-0.343,2.286c-0.907,0.905-2.255,1.438-3.786,1.438h-0.438v-8.04   c0.281-0.032,0.531-0.063,0.782-0.063C105.593,7.508,107.125,9.541,107.125,11.886 M105.904,4.474l-3.346,2.971V2.848   C103.748,3.317,104.812,3.88,105.904,4.474 M120.355,15.422l-0.188,0.124l-2.657-1.688V6.383l2.814,1.375L120.355,15.422z    M109.223,1.784l-0.283-0.25l-1.029,1.095c0,0-2.66-0.875-4.818-2.126c-2.971,2.251-5.223,3.846-8.131,5.911   c-0.689,1.031-1.064,2.376-1.064,3.849c0,4,3.691,7.129,8.353,7.129c3.032-1.281,5.19-3.064,6.787-5.224   c0.094-0.437,0.125-1.063,0.125-1.407c0-2.909-1.972-4.535-4.065-5.035L109.223,1.784z M122.671,6.632l-3.753-1.938   c-1.159,1.062-2.284,1.844-3.784,2.594c0.093,0.283,0.188,1.44,0.188,2.16v5.162c-0.346,0.373-0.69,0.562-1.033,0.562   c-0.344,0-0.719-0.158-1.189-0.501V2.16c0-0.72,0.438-1.565,1.004-1.565V0.503c-1.158,0-2.191,0.53-2.502,1.406   c-0.283-0.469-0.721-0.875-1.377-0.875l-0.031,0.126c0.375,0,0.75,1.062,0.75,2.062v10.042c0,0.969-0.127,1.471-0.5,2.002   l2.502,1.813l2.345-2.033l3.63,2.188c1.375-1.281,2.189-2.032,3.721-2.938c-0.093-0.471-0.125-1.377-0.125-1.752V8.103   C122.514,7.352,122.546,6.945,122.671,6.632 M63.275,15.671c-0.908-0.812-2.158-1.499-4.16-1.499l3.065-3.254l1.062,0.937   L63.275,15.671z M64.307,8.134l-2.626,2.846l-1.157-0.845V6.351C61.712,7.164,62.398,7.602,64.307,8.134 M77.225,15.733   L77.1,15.829l-2.502-1.533v-7.82l2.627,1.564V15.733z M79.663,14.576c-0.124-0.404-0.155-0.813-0.155-1.188V6.82L76.1,4.911   c-1.096,1.095-2.349,1.876-3.94,2.753c0.093,0.312,0.155,0.812,0.155,1.157v5.536c-0.5,0.533-1.002,0.846-1.438,0.846   c-0.437,0-0.876-0.188-1.346-0.688c0-0.438,0.032-5.192,0.032-8.101H72V5.819l-2.438-0.093c0,0,0.031-2.283,0.062-2.753h-0.281   l-3.41,3.597c-1.47-0.219-2.564-0.75-4.16-1.844c-1.25,1.281-2.096,1.875-3.596,2.595c0.062,0.751,0.093,1.157,0.093,1.909   c0,0.875,0,1.158-0.062,1.91l2.002,1.405l-1.877,1.972c-0.812,0.844-1.125,1.531-1.125,2.219c0,0.627,0.25,1.125,0.625,1.597   l0.22-0.124c-0.25-0.377-0.407-0.814-0.407-1.221c0-0.656,0.344-1.221,1.314-1.221c0.97,0,2.033,0.625,3.097,1.469   c1.188-1.094,2.065-1.812,3.377-2.469v-3.19c0-0.376,0.032-0.532,0.125-0.782L63.556,9.45l2.566-2.785h1.188   c0,2.846-0.031,6.506-0.031,6.912c0,0.501-0.094,1.127-0.437,1.627l2.846,2.095l2.44-2.22l3.534,2.312   C77.038,16.108,78.132,15.267,79.663,14.576 M91.018,15.015L90.8,14.767l-1.065,0.904c-0.529,0-0.905-0.374-0.905-0.969   c0-0.53,0.032-6.725,0.063-7.788c-0.408-0.815-0.939-1.439-2.129-1.846l-2.847,2.095c0-0.875-0.5-1.657-1.47-2.158l-1.658,1.409   l0.877,0.938c0,0-0.031,5.504-0.031,6.411c0,0.909-0.656,1.565-0.656,1.565l2.597,1.905l1.377-1.688l-1.034-0.938V7.727   c0.501-0.656,1.034-0.938,1.503-0.938c0.531,0,0.97,0.375,1.095,0.908c0.03,0.406,0.03,1.062,0.03,1.751   c0,1.062,0,4.814-0.03,6.285c0.188,0.658,1.156,1.563,1.782,1.563L91.018,15.015z M45.227,10.354l-0.062,4.502   c-0.468,0.252-0.781,0.408-1.219,0.72c-0.343-0.218-0.688-0.438-1.095-0.687V9.072L45.227,10.354z M45.572,6.445   c0,0.687-0.625,1.032-2.722,2.345V6.788c0-1,0.063-1.563,0.094-1.97h-0.437c0.062,0.438,0.094,0.97,0.094,1.97v7.979   c-1.313-0.783-2.909-1.598-4.599-1.723c1.596-0.906,3.19-2.126,3.19-3.91V5.225c0-1.25,0.5-2.003,1.815-3.127   c1.031,1.532,1.439,2.188,2.251,3.566C45.446,5.975,45.572,6.289,45.572,6.445 M54.423,15.671l-0.125,0.096l-2.471-1.504V6.351   l2.596,1.47V15.671z M41.475,1.565l-0.062-0.188c-1.314,0-2.033-0.157-2.722-0.281c-0.657-0.094-1.22-0.22-2.001-0.22   c-2.064,0-3.409,1.064-3.409,2.785c0,0.719,0.156,1.125,0.625,1.657l0.188-0.187c-0.25-0.313-0.375-0.626-0.375-0.971   c0-0.75,0.625-1.407,2.001-1.407c1.032,0,1.97,0.188,2.784,0.375c-1.751,1.281-2.816,2.22-2.816,4.941   c-0.25-0.094-0.594-0.156-1.063-0.156c-1.313,0-2.001,0.875-2.001,1.719c0,0.377,0.093,0.783,0.344,1.064l0.219-0.125   c-0.094-0.188-0.157-0.377-0.157-0.562c0-0.563,0.438-1.002,1.251-1.002c0.72,0,1.251,0.312,1.501,0.752   c0,1.906-0.375,2.877-1.375,2.877v0.219c1.501,0,3.314-1.095,3.314-2.846V6.507C37.721,3.629,39.41,2.286,41.475,1.565    M48.012,9.479l-2.785-1.845c1.064-0.657,1.408-0.845,1.783-1.126c0.658-0.47,0.845-0.751,0.845-1.032   c0-0.219-0.094-0.625-0.594-1.376c-0.625-0.938-1.22-1.846-2.471-3.596c-1.751,1.188-2.721,1.906-4.566,3.346   c-0.75,0.594-1.314,1.532-1.314,2.721c0,0.407,0.032,2.502,0.032,3.879c0,1.062-0.188,1.627-1.314,2.564   c-2.721,0-4.348,1.752-4.348,3.502c0,0.752,0.25,1.533,0.813,2.158l0.188-0.188c-0.345-0.377-0.625-0.971-0.625-1.657   c0-0.938,0.688-1.972,2.846-1.972c2.377,0,4.597,1.627,5.88,2.408c1.845-1.562,3.003-2.221,5.38-3.004v-3.972   C47.761,10.011,47.855,9.729,48.012,9.479 M56.957,6.757L53.36,4.786c-1.096,1.096-2.502,1.971-4.097,2.847   c0.219,0.657,0.155,1.5,0.155,2.564c0,2.598,0.031,4.005-0.438,4.692l3.971,2.439c1.096-1.128,2.284-2.002,3.88-2.878   c-0.125-0.407-0.157-0.813-0.157-1.188v-4.44C56.675,7.508,56.8,7.039,56.957,6.757 M4.473,11.638c0,0.938-0.157,1.502-0.657,1.938   C2.66,12.388,2.033,10.698,2.033,8.79c0-1.063,0.188-1.875,0.563-2.439c0.626-0.563,1.377-1.095,1.877-1.376v2.409V11.638z    M8.039,15.576c-1.563-0.187-2.909-0.781-3.91-1.688c1.939-1.406,2.377-2.222,2.377-3.785V3.754l0.062-0.062   c0.626,0.062,0.971,0.094,1.471,0.125V15.576z M26.805,8.915l-2.533,2.033V6.569l0.094-0.094L26.805,8.915z M28.775,13.952   l-0.188-0.189c-1.001,1.065-1.438,1.346-2.189,1.346c-0.562,0-1.22-0.342-2.126-1.188v-2.564l4.503-3.439L25.71,4.88   c-1.345,1.063-1.563,1.251-2.533,2.002c-0.532,0.405-0.845,0.594-1.189,0.75c0.062,0.375,0.125,0.907,0.125,1.282v5.726   c-0.531,0.498-1.032,0.719-1.407,0.719c-0.501,0-0.876-0.312-0.876-0.844V6.851c-0.312-0.969-0.906-1.595-2.064-2.001l-2.847,2.314   v-5.38c0-0.781,0.438-1.5,0.844-1.688V0c-1.157,0-1.97,0.752-2.314,1.628l-0.5-1.033l-0.188,0.032V14.39   c-1.125,0.594-2.345,1.25-3.691,1.25c-0.249,0-0.469,0-0.719-0.029V3.849c0.595,0.062,0.97,0.093,1.314,0.093   c0.531,0,0.876-0.093,1.126-0.219l1.533-3.128l-0.25-0.062c-0.594,1.188-0.906,1.375-2.189,1.375c-0.594,0-1.908-0.155-3.222-0.312   C5.411,1.44,4.128,1.284,3.473,1.284c-1.939,0-3.066,1.345-3.066,2.722c0,0.562,0.187,1.126,0.564,1.626l0.217-0.156   c-0.249-0.313-0.405-0.72-0.405-1.064c0-0.562,0.438-1.062,1.626-1.062c0.47,0,2.127,0.157,3.69,0.28   C4.316,4.505,2.847,5.443,1.126,6.757C0.375,7.758,0,9.041,0,10.354c0,3.377,2.565,7.035,8.227,7.035   c1.657-0.937,4.097-2.345,4.097-2.345l2.252,2.033l1.563-1.533l-1.22-1.282V7.727c0.532-0.719,1.126-1.032,1.596-1.032   c0.531,0,1.031,0.376,1.157,1.064v7.852c0.187,0.656,0.968,1.279,1.751,1.5l2.784-2.096l3.222,2.221L28.775,13.952z"></path>

	<!-- fallback .png -->
	<image src="/bg-images/png/bg-logo--full.png" xlink:href=""/>
</svg>
</span>
			<span class="site-header__bug-link">
				<svg width="20" height="20" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="globe-title">
	<title id="globe-title">Visit The Boston Globe</title>
	<path d="M15.314 10.916l-.068 4.873c-.507.27-.845.44-1.32.777-.372-.236-.744-.474-1.184-.745V9.528l2.572 1.388zm.373-4.232c0 .744-.677 1.117-2.945 2.54v-2.17c0-1.082.068-1.69.1-2.132h-.472c.067.474.102 1.05.102 2.133v8.633c-1.423-.848-3.15-1.728-4.977-1.863 1.727-.982 3.452-2.302 3.452-4.233v-4.23c0-1.354.542-2.167 1.964-3.385 1.12 1.657 1.56 2.368 2.44 3.858.203.338.34.677.34.846l-.003.004zm-4.434-5.28l-.068-.204c-1.422 0-2.2-.17-2.946-.304C7.53.794 6.92.658 6.072.658c-2.235 0-3.69 1.152-3.69 3.013 0 .78.168 1.22.676 1.797l.202-.203c-.27-.34-.405-.678-.405-1.05 0-.812.676-1.523 2.165-1.523 1.118 0 2.133.204 3.013.407C6.14 4.483 4.99 5.5 4.99 8.444c-.272-.102-.644-.17-1.152-.17-1.42 0-2.166.95-2.166 1.862 0 .407.1.848.373 1.152l.237-.135c-.102-.204-.17-.407-.17-.61 0-.61.474-1.083 1.354-1.083.78 0 1.354.338 1.624.813 0 2.064-.405 3.114-1.488 3.114v.236c1.624 0 3.588-1.184 3.588-3.08v-3.79c0-3.116 1.828-4.57 4.063-5.35zm7.075 8.563L15.314 7.97c1.15-.71 1.524-.914 1.93-1.218.71-.508.914-.813.914-1.117 0-.238-.1-.677-.643-1.49-.677-1.016-1.32-1.997-2.674-3.892-1.894 1.286-2.943 2.063-4.94 3.622C9.09 4.517 8.48 5.533 8.48 6.82c0 .44.034 2.707.034 4.196 0 1.15-.202 1.76-1.42 2.775-2.945 0-4.706 1.898-4.706 3.793 0 .813.27 1.66.88 2.335l.204-.204c-.373-.407-.677-1.05-.677-1.793 0-1.016.746-2.133 3.08-2.133 2.574 0 4.977 1.76 6.365 2.607 1.997-1.693 3.25-2.404 5.822-3.252v-4.298c0-.304.1-.608.27-.88l-.002.003z" fill-rule="nonzero" fill="#000"/>

	<!-- fallback .png -->
	<image src="/bg-images/png/bg-logo--bug.png" xlink:href=""/>
</svg>
</span>
		</a>
	</div>

	<div class="js-user-settings" data-location="" data-baseurl="https://www.bostonglobe.com"></div>
		<div class="js-user-subscribe" data-location="" data-baseurl="https://www.bostonglobe.com"></div>
	<!-- Current section link. -->
	<a href="/news/bigpicture?p1=BGHeader_SectionLink" class="site-header__section-front site-header__section-front--long" aria-label="current sectionfront The Big Picture">The Big Picture</a>
      	<div class="smartbar ">
		<div class="smartbar__slot smartbar__slot--scoreboard  scoreboard ">

			<!-- Active game scores get appended here  -->
			<ul class="scoreboard__list js-sport-scores" data-order="nhl"></ul>

			<!-- "Next Score" button. -->
			<div class="scoreboard__next-score js-next-score">
				<span class="scoreboard__next-score-text" onclick="var s=s_gi('nytbostonglobecom');	s.tl(this,'o','BG Header – Next Score');">Next Score</span>
				<svg width="13" height="9" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="arrow-down-title">
	<title id="arrow-down-title">View the next score</title>
	<path d="M1.847.377L.14 2.024 6.333 8.45l6.195-6.425L10.82.377 6.334 5.03" fill="#777"/>

	<!-- fallback .png -->
	<image src="/bg-images/png/arrow-down.png" xlink:href=""/>
</svg>
</div>
		</div>
	</div>
<div class="site-header__centerpiece">
			<a href="/news/bigpicture?p1=BGHeader_SectionTitle" aria-label="current sectionfront The Big Picture" class="site-header__centerpiece-link site-header__centerpiece-link--long">The Big Picture</a>
						</div>

		<nav class="site-nav">
			<ul class="site-nav__list" role="menu">

				</ul>
		</nav>
	</header>
<!-- Page header ends -->

<script>
$( document ).ready(function() {
	$('.social-icon.email').removeAttr('onclick');
	$('.social-icon.print').removeAttr('onclick');
	$('.js-comments-toggle').removeAttr('onclick');
	$('.tools .bg-button').removeAttr('onclick');
	$('.comment-controls .btn.med').removeAttr('onclick');

	$('.js-comment-count-label').click(function(){
		console.log('open comments');
		console.log($('.comment-controls .btn').hasClass( "med" ));
	})



	$('.js-comments-toggle').click(function(){
		var s=s_gi('nytbostonglobecom');
		s.tl(this,'o','Gallery_Comments_Top');
	})

	$('.Top .social-icon.email').click(function(){
		var s=s_gi('nytbostonglobecom');
		s.tl(this,'o','GalleryTools_Email_Top');
	});

	$('.Top .social-icon.print').click(function(){
		var s=s_gi('nytbostonglobecom');
		s.tl(this,'o','GalleryTools_Print_Top');
	});

	$('.Bottom .social-icon.email').click(function(){
		var s=s_gi('nytbostonglobecom');
		s.tl(this,'o','GalleryTools_Email_Bottom');
	});

	$('.Bottom .social-icon.print').click(function(){
		var s=s_gi('nytbostonglobecom');
		s.tl(this,'o','GalleryTools_Print_Bottom');
	});

	$('.Top.tools .bg-button').click(function(){
		var s=s_gi('nytbostonglobecom');
		s.tl(this,'o','GalleryTools_Send_Top');
	});
	$('.Bottom.tools .bg-button').click(function(){
		var s=s_gi('nytbostonglobecom');
		s.tl(this,'o','GalleryTools_Send_Bottom');
	});
});
</script>


<div id="container">

	<section>
	  	<h3 class="pictureInfo-headline">Globe staff photos of the month, February 2019</h3>
	  	<div class="subhead geor"><!---->Here&rsquo;s a look at some of the best images taken by Globe photographers last month: winter weather, presidential candidates on the stump, a Chinese New Year celebration, Red Sox spring training, and the Patriots&rsquo; Super Bowl LIII victory parade.</div>

		<!-- position:Top -->
<!-- rewriteurl  | /Boston/2011-2020/2019/02/02/BostonGlobe.com/BigPicture/Galleries/19feb_pom.gallery.xml-->

<!--
METADATA FOR EMTAF
      <headline>Globe staff photos of the month, February 2019</headline>
      <source>
</source>
      <teasetext>Here&amp;rsquo;s a look at some of the best images taken by Globe photographers last month: winter weather, presidential candidates on the stump, a Chinese New Year celebration, Red Sox spring training, and the Patriots&amp;rsquo; Super Bowl LIII victory parade.</teasetext>
      <byline>
</byline>
      <date>20190302143356</date>
-->

<aside class="social-tools">
	<ul class="Top social-tools__list">
		<!-- imgAltCaption: Boston, MA - 2/05/2019 - New England Patriots wide receiver and Super Bowl LIII MVP Julian Edelman taps his heart and points to the fans to show his love during today's New England Patriots Super Bowl LIII Victory Parade. (Barry Chin/Globe Staff), Section: Metro, Reporter: Globe Staff, Topic: 06parade, LOID: 8.5.332310335 -->
		<li class="social-tools__list-item">
	      <a class="bg-button bg-button--article-tool js-email-button" onClick="var s=s_gi('nytbostonglobecom');s.tl(this,'o','ArticleTools_Email_Top');" data-omniture="BG Gallery Share Tools - Email - Top" title="Share via E-mail" aria-label="Share via E-mail" href="javascript:;">
	      	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="email-blue-icon"><title id="email-blue-icon">Email to a Friend</title><g fill="#127694" fill-rule="evenodd"><path d="M21.658 4.403H2.393c2.06 1.93 8.782 8.235 9.632 9.033l9.633-9.033M15.985 11.925l6.848 6.42V5.503"/><path d="M12.026 15.638L9.24 13.025l-6.846 6.422h19.264c-1.304-1.226-4.47-4.195-6.845-6.422l-2.787 2.613zM1.22 5.503v12.842l6.846-6.42"/></g>
	<!-- fallback .png -->
	<image src="/bg-images/png/email-blue-small.png" xlink:href=""/>
</svg>
</a>
	      <div class="bg-form social-tools__email-form js-email-box">
				<h4 class="bg-form__title">Share via e-mail</h4>
				<div class="bg-form__error js-email-error"></div>
				<form name="emailForm" action="" method="post" class="js-email-form">
					<input name="story_url" type="hidden" id="story_url-Top" class="js-story-url" value="https://www.bostonglobe.com/news/bigpicture/2019/03/02/globe-staff-photos-month-february/zsbKlFlxkRDAiLYNucIwxJ/story.html"/>
					<label for="email-to-primary-Top" class="bg-form__label">To</label>
					<input name="recipient_email" type="text" id="email-to-primary-Top" placeholder="Separate multiple addresses with a comma" class="bg-form__input js-email-to" />
					<label for="email-mssg-primary-Top" class="bg-form__label">Add a message</label>
					<textarea name="message" id="email-mssg-primary-Top" maxlength="80" placeholder="80 character limit" class="bg-form__textarea js-email-message"></textarea>
					<label for="email-from-primary-Top" class="bg-form__label">Your e-mail</label>
					<input name="sender_email" type="text" id="email-from-primary-Top" class="bg-form__input js-email-from" />
					<input type="submit" value="Send" class="bg-button bg-button--submit js-email-submit" onClick="var s=s_gi('nytbostonglobecom');s.tl(this,'o','ArticleTools_Send_Top');"/>
				</form>
			</div><!-- /end .email-form -->
	    </li>
	    <li class="social-tools__list-item">
	      <a class="bg-button social-tools__link social-tools__link--facebook" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Facebook');" href="http://www.facebook.com/sharer.php?u=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25&amp;t=Globe staff photos of the month, February 2019" target="_blank" title="Share on Facebook" aria-label="Share on Facebook">
	      	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="facebook-share-icon"><title id="facebook-share-icon">Share on Facebook</title><path d="M13.503 21.63v-8.785h2.97l.445-3.425h-3.415V7.234c0-.99.277-1.667 1.71-1.667h1.826V2.503c-.317-.042-1.402-.135-2.663-.135-2.634 0-4.437 1.595-4.437 4.525V9.42H6.96v3.425h2.98v8.786h3.563" fill="#FFF" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/facebook-small.png" xlink:href=""/>
</svg>
</a>
	    </li>
	    <li class="social-tools__list-item">
	      <a class="bg-button social-tools__link social-tools__link--twitter" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Twitter');" href="http://twitter.com/intent/tweet?text=Globe staff photos of the month, February 2019&amp;url=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25&amp;via=BostonGlobe" target="_blank" title="Share on Twitter" aria-label="Share on Twitter">
	      	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="twitter-share-icon"><title id="twitter-share-icon">Share on Twitter</title><path d="M22.457 5.614c-.77.336-1.596.563-2.464.665.886-.524 1.566-1.352 1.887-2.338-.83.484-1.747.836-2.725 1.025-.783-.82-1.898-1.334-3.132-1.334-2.37 0-4.29 1.89-4.29 4.224 0 .332.037.654.11.963-3.566-.176-6.728-1.858-8.845-4.414-.37.624-.58 1.35-.58 2.124 0 1.466.757 2.76 1.91 3.516-.705-.02-1.367-.212-1.945-.528v.053c0 2.048 1.48 3.755 3.442 4.144-.36.096-.74.148-1.13.148-.277 0-.546-.027-.808-.076.546 1.68 2.13 2.9 4.01 2.934-1.47 1.133-3.32 1.808-5.33 1.808-.347 0-.688-.02-1.024-.058 1.9 1.198 4.154 1.897 6.578 1.897 7.893 0 12.21-6.437 12.21-12.02 0-.183-.006-.365-.013-.546.838-.595 1.565-1.34 2.14-2.186" fill="#FFF" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/twitter-small.png" xlink:href=""/>
</svg>
</a>
	    </li class="social-tools__list-item">
	    <li class="social-tools__list-item">
			<a class="bg-button social-tools__link social-tools__link--google-plus" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Google+');" href="https://plus.google.com/share?url=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25%26s_campaign%3dsm_gp%26hl=en-US" target="_blank" title="Share on Google+" aria-label="Share on Google+">
				 		<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="google-share-icon"><title id="google-share-icon">Share on Google Plus</title><g fill="#FFF" fill-rule="evenodd"><path d="M7.82 10.955v2.508h4.15c-.17 1.076-1.255 3.156-4.15 3.156-2.497 0-4.535-2.07-4.535-4.62 0-2.55 2.038-4.62 4.535-4.62 1.42 0 2.372.607 2.916 1.13l1.985-1.912c-1.274-1.192-2.925-1.913-4.9-1.913C3.776 4.685.505 7.955.505 12c0 4.044 3.27 7.315 7.315 7.315 4.222 0 7.022-2.968 7.022-7.148 0-.48-.052-.846-.115-1.212H7.82zM23.495 10.955h-2.09v-2.09h-2.09v2.09h-2.09v2.09h2.09v2.09h2.09v-2.09h2.09"/></g>
	<!-- fallback .png -->
	<image src="/bg-images/png/google-plus-small.png" xlink:href=""/>
</svg>
</a>
				</li>
	    <li class="social-tools__list-item">
	      <a class="bg-button social-tools__link social-tools__link--linkedin" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Linkedin');" href="http://www.linkedin.com/shareArticle?mini=true&amp;url=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25&amp;title=Globe staff photos of the month, February 2019&amp;summary=Here%26amp%3Brsquo%3Bs+a+look+at+some+of+the+best+images+taken+by+Globe+photographers+last+month%3A+winter+weather%2C+presidential+candidates+on+the+stump%2C+a+Chinese+New+Year+celebration%2C+Red+Sox+spring+training%2C+and+the+Patriots%26amp%3Brsquo%3B+Super+Bowl+LIII+victory+parade." target="_blank" title="Share on LinkedIn" aria-label="Share on LinkedIn"><svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="linkedin-share-icon"><title id="linkedin-share-icon">Share on LinkedIn</title><g fill="#FFF" fill-rule="evenodd"><path d="M5.12 2.96c1.2 0 2.177.974 2.177 2.176 0 1.202-.976 2.177-2.176 2.177-1.205 0-2.177-.975-2.177-2.177S3.915 2.96 5.12 2.96zM3.24 21.04H7V8.964H3.24V21.04zM9.353 8.963h3.598v1.65h.053c.5-.948 1.726-1.95 3.55-1.95 3.802 0 4.504 2.5 4.504 5.754v6.624h-3.752v-5.873c0-1.4-.025-3.202-1.95-3.202-1.954 0-2.252 1.526-2.252 3.102v5.974h-3.75V8.964z"/></g>
	<!-- fallback .png -->
	<image src="/bg-images/png/linkedin-small.png" xlink:href=""/>
</svg>
</a>
	    </li>
	</ul>
	<ul class="social-tools__list social-tools__list--right">
		<li class="social-tools__list-item social-tools__list-item--print">
			<a class="bg-button bg-button--article-tool" onclick="var s=s_gi('nytbostonglobecom');s.tl(this,'o','ArticleTools_Print_Top')" href="javascript:print();">
				<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="print-blue-icon"><title id="print-blue-icon">Print this Article</title><path d="M17.935 12.547h4.388v9.755H1.677v-9.755H6.01V1.697h11.925v10.85zM8.047 6.92v1.477h7.906V6.92H8.047zm0-2.96v1.476h7.906V3.96H8.047zm0 5.922v1.475h7.906V9.882H8.047zm-3.12 5.892v3.288H19.01v-3.288H4.927z" fill="#127694" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/print-blue-small.png" xlink:href=""/>
</svg>
<span class="hidden">Print</span>
			</a>
		</li>
		<li class="social-tools__list-item">
		<a href="#comments" class="js-comments-toggle bg-button bg-button--article-tool">
			  <svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="comments-blue-icon"><title id="comments-blue-icon">View Comments</title><path d="M2.576 1.63h18.848v15.08H2.576V1.63zm1.877 15.08h5.66l-5.66 5.66v-5.66z" fill="#127694" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/comment-blue-small.png" xlink:href=""/>
</svg>
<span class="js-comment-count social-tools__comment-count"></span>
			  <span class="js-comment-count-label hidden">Comments</span>
			</a>
		</li>
	</ul>
	</aside>
<div class="sticky-tools">
		<a class="sticky-tools__button sticky-tools__button--logo" href="/?p1=BGHeader_Logo_Sticky">
			<svg width="30" height="35" viewBox="0 0 30 35" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="globe-medium-title">
	<title id="globe-medium-title">Visit The Boston Globe</title>
	<path d="M24.288 18.943l-.12 8.657c-.903.484-1.505.785-2.35 1.384-.662-.418-1.325-.84-2.11-1.322V16.477l4.58 2.466zm.665-7.517c0 1.32-1.205 1.984-5.244 4.51v-3.85c0-1.924.12-3.005.18-3.79h-.843c.12.843.18 1.866.18 3.79v15.338c-2.53-1.505-5.605-3.07-8.86-3.31 3.074-1.743 6.147-4.088 6.147-7.52V9.08c0-2.406.965-3.85 3.497-6.015C22 6.01 22.786 7.275 24.35 9.92c.362.603.603 1.205.603 1.506zm-7.895-9.382l-.12-.362c-2.533 0-3.92-.302-5.246-.54C10.427.962 9.342.72 7.836.72c-3.98 0-6.57 2.046-6.57 5.353 0 1.384.3 2.164 1.204 3.187l.362-.36c-.483-.603-.722-1.203-.722-1.866 0-1.442 1.205-2.705 3.856-2.705 1.99 0 3.798.36 5.365.72-3.374 2.465-5.425 4.27-5.425 9.502-.482-.18-1.145-.3-2.05-.3C1.326 14.253 0 15.936 0 17.56c0 .724.18 1.506.663 2.046l.422-.24c-.182-.362-.303-.722-.303-1.082 0-1.082.844-1.926 2.412-1.926 1.386 0 2.41.602 2.89 1.446 0 3.667-.72 5.532-2.65 5.532v.42c2.893 0 6.39-2.105 6.39-5.472v-6.736c0-5.534 3.255-8.118 7.234-9.502zm12.597 15.214L24.29 13.71c2.05-1.262 2.713-1.624 3.435-2.164 1.267-.903 1.63-1.444 1.63-1.985 0-.42-.183-1.202-1.147-2.645-1.205-1.805-2.35-3.55-4.76-6.915-3.377 2.285-5.244 3.665-8.8 6.434-1.447 1.14-2.533 2.946-2.533 5.23 0 .784.06 4.813.06 7.46 0 2.043-.36 3.128-2.53 4.93-5.243 0-8.378 3.37-8.378 6.736 0 1.444.483 2.948 1.567 4.148l.362-.36c-.663-.723-1.204-1.864-1.204-3.186 0-1.806 1.325-3.792 5.484-3.792 4.58 0 8.858 3.13 11.33 4.634 3.556-3.008 5.787-4.27 10.367-5.776V18.82c0-.538.18-1.08.482-1.562z" fill="#231F20" fill-rule="evenodd"/>

	<!-- fallback .png -->
	<image src="/bg-images/png/bg-logo--bug--medium.png" xlink:href=""/>
</svg>
</a>
		<a class="sticky-tools__button sticky-tools__button--twitter" data-omniture="BG Article Share Tools - Twitter - Sticky" href="http://twitter.com/intent/tweet?text=Globe staff photos of the month, February 2019&amp;url=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25%3Fevent%3Devent25&amp;via=BostonGlobe" target="_blank" title="Share on Twitter">
			<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="twitter-share-icon"><title id="twitter-share-icon">Share on Twitter</title><path d="M22.457 5.614c-.77.336-1.596.563-2.464.665.886-.524 1.566-1.352 1.887-2.338-.83.484-1.747.836-2.725 1.025-.783-.82-1.898-1.334-3.132-1.334-2.37 0-4.29 1.89-4.29 4.224 0 .332.037.654.11.963-3.566-.176-6.728-1.858-8.845-4.414-.37.624-.58 1.35-.58 2.124 0 1.466.757 2.76 1.91 3.516-.705-.02-1.367-.212-1.945-.528v.053c0 2.048 1.48 3.755 3.442 4.144-.36.096-.74.148-1.13.148-.277 0-.546-.027-.808-.076.546 1.68 2.13 2.9 4.01 2.934-1.47 1.133-3.32 1.808-5.33 1.808-.347 0-.688-.02-1.024-.058 1.9 1.198 4.154 1.897 6.578 1.897 7.893 0 12.21-6.437 12.21-12.02 0-.183-.006-.365-.013-.546.838-.595 1.565-1.34 2.14-2.186" fill="#FFF" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/twitter-small.png" xlink:href=""/>
</svg>
</a>
		<a class="sticky-tools__button sticky-tools__button--facebook" data-omniture="BG Article Share Tools - Facebook - Sticky" href="http://www.facebook.com/sharer.php?u=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25%3Fevent%3Devent25&amp;tGlobe staff photos of the month, February 2019" target="_blank" title="Share on Facebook">
			<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="facebook-share-icon"><title id="facebook-share-icon">Share on Facebook</title><path d="M13.503 21.63v-8.785h2.97l.445-3.425h-3.415V7.234c0-.99.277-1.667 1.71-1.667h1.826V2.503c-.317-.042-1.402-.135-2.663-.135-2.634 0-4.437 1.595-4.437 4.525V9.42H6.96v3.425h2.98v8.786h3.563" fill="#FFF" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/facebook-small.png" xlink:href=""/>
</svg>
</a>
		<a class="sticky-tools__button sticky-tools__button--comment js-sticky-comments" href="#comments" onClick="var s=s_gi('nytbostonglobecom');s.trackLinkVars='eVar15,channel,prop1';s.trackLinkEvents='none';s.tl(this,'o', 'Article_Comments - Sticky');">
			<svg version="1.1" id="comment-icon" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="15px" height="17.063px" viewBox="0 0 15 17.063" enable-background="new 0 0 15 17.063" xml:space="preserve" aria-labelledby="comment-title">
	<title id='comment-title'>Comment on this</title>
	<polygon id="path-1" points="2.118,17.063 2.118,12.461 0,12.461 0,0 15,0 15,12.461 7.588,12.461 "/>

	<!-- fallback .png -->
	<image src="/bg-images/png/comment.png" xlink:href=""/>
</svg>
<span class="js-comment-count"></span>
		</a>
		<a class="sticky-tools__button sticky-tools__button--top js-back-to-top" href="javascript:;" onClick="var s=s_gi('nytbostonglobecom');s.trackLinkVars='none';s.trackLinkEvents='none';s.tl(this,'o', 'Article_Back to Top - Sticky');">
			<svg width="17" height="11" viewBox="0 0 17 11" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="arrow-up-title">
	<title id="arrow-up-title">Scroll to top of page</title>
	<path d="M8.633.574L8.5.698 8.365.574.603 7.744l2.903 2.682L8.5 5.814l4.994 4.612 2.903-2.68" fill="#000" fill-rule="evenodd"/>

	<!-- fallback .png -->
	<image src="/bg-images/png/arrow-up.png" xlink:href=""/>
</svg>
</a>
	</div>
<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/14/BostonGlobe.com/Metro/Images/Rinaldi14mccue01.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">1</div>
					<div class="gcaption geor"><!---->Iraq war veteran Peter Rooney placed his hand on the casket of WWII veteran James E. McCue following a graveside military service for McCue who died at the age of 97. (Jessica Rinaldi/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/21/BostonGlobe.com/Metro/Images/tlumackirally299.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">2</div>
					<div class="gcaption geor"><!---->Tina Gweah, a senior at Chelsea High School, joined several hundred youths from across the state for a on Boston Common Bandstand to call for more state funding for schools, job training, and juvenile justice reforms. Representative Ayanna Pressley spoke to the crowd before the group marched to the State House. (John Tlumacki/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/21/BostonGlobe.com/Metro/Images/ryanice1met.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">3</div>
					<div class="gcaption geor"><!---->MIT boats, feet, shovels, and poles were used to break up ice on the Charles River to get sailing season underway. (David L. Ryan/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/10/BostonGlobe.com/Regional/Images/klima_17soquincy_20190210_007.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">4</div>
					<div class="gcaption geor"><!---->Carrie Deng (center left) and Ritage Elhachimi (center right) played together before performing at the 31st annual Lunar New Year Festival at North Quincy High School on Feb. 10. (Nathan Klima for The Boston Globe) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/26/BostonGlobe.com/Metro/Images/turner022619METROfeature78.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">5</div>
					<div class="gcaption geor"><!---->A school bus was reflected in patterned steel wrapping a South Boston diner at West Broadway and A Street. (Lane Turner/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/27/BostonGlobe.com/Metro/Images/turner022719METRO28hampton276.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">6</div>
					<div class="gcaption geor"><!---->Hampton, NH - 02/26/19 - Sharon Morey (cq) of Live and Let Live Farm Sharon Morey of Live and Let Live Farm Rescue  Sanctuary in Chichester, N.H., held one of 16 guinea pigs rescued by a neighbor from house fire in Hampton, N.H. A young boy was killed in the blaze. (Lane Turner/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/24/BostonGlobe.com/Metro/Images/Rinaldi19224Furries02.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">7</div>
					<div class="gcaption geor"><!---->An attendee at Anthro New England took a break during an anthropromorphism and furry convention at the Boston Park Plaza on Feb. 24. (Jessica Rinaldi/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/23/BostonGlobe.com/Metro/Images/lee022319mansfieldlplanecrash4MET.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">8</div>
					<div class="gcaption geor"><!---->Emergency personnel worked at the scene of a single-engine plane crash at the Mansfield Municipal Airport. (Matthew J. Lee/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/19/BostonGlobe.com/Metro/Images/klima_18holycross_20190218_001.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">9</div>
					<div class="gcaption geor"><!---->Students (from left) Carson Harold, Mithra Salmassi, and Caroline O&rsquo;Conner posed with signs they held during a sit-in at the College of the Holy Cross. (Nathan Klima for The Boston Globe) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/16/BostonGlobe.com/Metro/Images/Walker_021619_17NewHampshire_12319x.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">10</div>
					<div class="gcaption geor"><!---->Senator Cory Booker of New Jersey shared a laugh with supporters while signing an autograph during a campaign event at 3S Art Space in Portsmouth, N.H. (Craig F. Walker/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/09/BostonGlobe.com/Metro/Images/Walker_020819_09warren_4657x.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">11</div>
					<div class="gcaption geor"><!---->Senator Elizabeth Warren officially announced her presidential run at Everett Mills in Lawrence. (Craig F. Walker/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/19/BostonGlobe.com/Metro/Images/Walker_021919_20harris_15158x.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">12</div>
					<div class="gcaption geor"><!---->Senator Kamala Harris spoke with an attendee during &ldquo;Politics  Eggs&rdquo; at the New Hampshire Institute of Politics at St. Anselm College in Manchester, N.H. (Craig F. Walker/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/18/BostonGlobe.com/Metro/Images/RinaldiPrezDay02.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">13</div>
					<div class="gcaption geor"><!---->Elliott Schmitt, 7, of Portland, Maine, posed for a picture with President John Adams, played by Michael LePage, during a Presidents Day celebration at the JFK Library. The Schmitt family goes all out celebrating Presidents Day. The kids set out their shoes the night before and a &ldquo;ghost of a former president&rdquo; fills them with candy and marks a spot in the Constitution for the family to read. (Jessica Rinaldi/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/01/BostonGlobe.com/Metro/Images/Walker_013019_xxpiano_5658x.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">14</div>
					<div class="gcaption geor"><!---->Sonya Bandouil worked on her balance with Dr. Rachel Benjamin at Spaulding Rehabilitation Center, while her boyfriend, Alex Pankiewicz, watched. Bandouil, a pianist, was critically injured when the fa&ccedil;ade of an Allston restaurant collapsed and fell on her. Her hand was crushed, and she lost a finger, but has returned to playing the piano. (Craig F. Walker/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/04/BostonGlobe.com/Metro/Images/KREITER02042019pierspark4.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">15</div>
					<div class="gcaption geor"><!---->Josette Roysin, 3, played at Piers Park in East Boston with her mother, Stacy. Josette came to Boston from North Carolina for a heart transplant at Boston Children&rsquo;s Hospital three months ago and is here for a bit longer during her follow up care. (Suzanne Kreiter/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/08/BostonGlobe.com/Lifestyle/Images/Walker_020819_04WalshOpioidspic_3402x.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">16</div>
					<div class="gcaption geor"><!---->Mayor Martin J. Walsh posed for a portrait at City Hall. (Craig F. Walker/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/17/BostonGlobe.com/Metro/Images/RinaldiCNY12.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">17</div>
					<div class="gcaption geor"><!---->A lion dancer passed a woman during the annual Chinese New Year celebration and parade in Boston. (Jessica Rinaldi/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/18/BostonGlobe.com/Metro/Images/turner013019LIVToddGieg51[1].jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">18</div>
					<div class="gcaption geor"><!---->Todd Gieg of Lynn is building a model of the Boston-Revere-Lynn narrow gauge railroad and the 1895 world it traversed. (Lane Turner/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/07/BostonGlobe.com/Metro/Images/Walker_020719_ballerinaspic_2383x.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">19</div>
					<div class="gcaption geor"><!---->Kendall Case, 11, stretched while preparing to perform in the Pre-Competitive Classical Competition of the Youth America Grand Prix at John Hancock Hall. Hundreds of ballet dancers from all over New England, ages 9 to 19, auditioned for the Youth America Grand Prix. (Craig F. Walker/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/06/BostonGlobe.com/Arts/Images/RinaldiART02-001.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">20</div>
					<div class="gcaption geor"><!---->Actor Jo Yang rehearsed in a pool for the ART&rsquo;s &ldquo;Endlings.&rdquo; (Jessica Rinaldi/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/07/BostonGlobe.com/Metro/Images/Rinaldi09dudley03.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">21</div>
					<div class="gcaption geor"><!---->Ivonne Bonilla of Roxbury and her dog, Chipo, walked past empty storefronts along Washington Street in Dudley Square. (Jessica Rinaldi/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/08/BostonGlobe.com/Metro/Images/Walker_013119_03Farragherxx_7130x[1][1].jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">22</div>
					<div class="gcaption geor"><!---->Glenda Savitz signed to her 2-year-old daughter, Samantha, at their home in Auburndale. Some twenty neighbors have enrolled in a sign language class so they can communicate with Samantha, who is deaf. (Craig F. Walker/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/14/BostonGlobe.com/Metro/Images/wiggs_Valentine_01.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">23</div>
					<div class="gcaption geor"><!---->Steven Wang and Montserrat Tello danced at the Senior Citizen Valentine&rsquo;s Day Dance sponsored by the Boston Red Sox at the State Street Pavilion inside Fenway Park. (Jonathan Wiggs/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/08/BostonGlobe.com/Metro/Images/KREITER02082019steamandshadow1.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">24</div>
					<div class="gcaption geor"><!---->A pedestrian seemed to emerge from a cloud escaping from a manhole cover on High Street in downtown Boston. (Suzanne Kreiter/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/14/BostonGlobe.com/Sports/Images/davisbosdet8spts.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">25</div>
					<div class="gcaption geor"><!---->The Celtics Marcus Smart celebrated with the crowd after he hit an unlikely shot while hitting the floor in the third quarter that gave Boston a 91-72 lead over the Detroit Pistons. (Jim Davis/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/19/BostonGlobe.com/Sports/Images/davismtn1Aspts.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">26</div>
					<div class="gcaption geor"><!---->Methuen-Tewksbury&rsquo;s Cassidy Gruning leaped in celebration after she scored a first-period goal to put her team ahead of Newburyport, 1-0. (Jim Davis/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/12/BostonGlobe.com/Sports/Images/davisnubc10spts.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">27</div>
					<div class="gcaption geor"><!---->Northeastern players celebrated with goalie Cayden Primeau after the final horn sounded in the Huskies&rsquo; Beanpot victory.) (Jim Davis/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/13/BostonGlobe.com/Sports/Images/lee021219buharvard15SPTS.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">28</div>
					<div class="gcaption geor"><!---->Cambridge MA 2/12/19 Boston University Boston University&rsquo;s Abbey Stanley (21) and her teammates celebrated with the Beanpot trophy after they defeated Harvard 3-2 during overtime. (Matthew J. Lee/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/20/BostonGlobe.com/ReceivedContent/Images/post8.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">29</div>
					<div class="gcaption geor"><!---->Red Sox infielder Tzu-Wei Lin took a swing at JetBlue Park. (Stan Grossfeld/ Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/23/BostonGlobe.com/ReceivedContent/Images/sox6.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">30</div>
					<div class="gcaption geor"><!---->Yankees pitcher Rex Brothers arrived at JetBlue Park for a spring training game. (Stan Grossfeld/ Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/19/BostonGlobe.com/ReceivedContent/Images/Chin021919Day9RedSoxSpringTraining_Spt1.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">31</div>
					<div class="gcaption geor"><!---->Boston Red Sox second baseman Brock Holt posed for a photo during Picture Day at JetBlue Park. (Barry Chin/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/04/BostonGlobe.com/ReceivedContent/Images/stansb11.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">32</div>
					<div class="gcaption geor"><!---->Rob Gronkowski made a key reception late during the fourth quarter of the Super Bowl. (Stan Grossfeld/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/04/BostonGlobe.com/ReceivedContent/Images/jimsb26.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">33</div>
					<div class="gcaption geor"><!---->David Andrews and other Patriots celebrated a TD during the fourth quarter of the Super Bowl. (Jim Davis/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/04/BostonGlobe.com/Sports/Images/Chin020319Patriots-RamsSuperBowlLIII_Spt8.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">34</div>
					<div class="gcaption geor"><!---->New England Patriots quarterback Tom Brady and his daughter Vivian celebrated the Super Bowl win. (Barry Chin/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/04/BostonGlobe.com/Sports/Images/davissbfollow8Aspts.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">35</div>
					<div class="gcaption geor"><!---->J.C. Jackson leaped into the arms of teammate Dont&rsquo;a Hightower as the final seconds ticked off the clock in New England&rsquo;s Super Bowl victory. (Jim Davis/Globe Staff) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/05/BostonGlobe.com/Metro/Images/klima_06parade_20190205_012.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">36</div>
					<div class="gcaption geor"><!---->Tom Brady held up the Super Bowl trophy during the Super Bowl parade through Boston. (Nathan Klima for The Boston Globe) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/05/BostonGlobe.com/Metro/Images/klima_06parade_20190205_008.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">37</div>
					<div class="gcaption geor"><!---->Patriots fans celebrated during the Super Bowl parade. (Nathan Klima for The Boston Globe) </div>
					</article>
				<div class="photo"><img src="//c.o0bg.com/rf/image_1200w/Boston/2011-2020/2019/02/05/BostonGlobe.com/Metro/Images/Chin020519PatriotsParade_Spt3.jpg"/></div>
			    <article class="pcaption">
					<div class="bignum">38</div>
					<div class="gcaption geor"><!---->New England Patriots wide receiver and Super Bowl LIII MVP Julian Edelman tapped his heart and pointed to the fans to show his love during the victory parade. (Barry Chin/Globe Staff) </div>
					</article>
				</section>
</div>

<!-- position:Bottom -->
<!-- rewriteurl  | /Boston/2011-2020/2019/02/02/BostonGlobe.com/BigPicture/Galleries/19feb_pom.gallery.xml-->

<!--
METADATA FOR EMTAF
      <headline>Globe staff photos of the month, February 2019</headline>
      <source>
</source>
      <teasetext>Here&amp;rsquo;s a look at some of the best images taken by Globe photographers last month: winter weather, presidential candidates on the stump, a Chinese New Year celebration, Red Sox spring training, and the Patriots&amp;rsquo; Super Bowl LIII victory parade.</teasetext>
      <byline>
</byline>
      <date>20190302143356</date>
-->

<aside class="social-tools">
	<ul class="Bottom social-tools__list">
		<!-- imgAltCaption: Boston, MA - 2/05/2019 - New England Patriots wide receiver and Super Bowl LIII MVP Julian Edelman taps his heart and points to the fans to show his love during today's New England Patriots Super Bowl LIII Victory Parade. (Barry Chin/Globe Staff), Section: Metro, Reporter: Globe Staff, Topic: 06parade, LOID: 8.5.332310335 -->
		<li class="social-tools__list-item">
	      <a class="bg-button bg-button--article-tool js-email-button" onClick="var s=s_gi('nytbostonglobecom');s.tl(this,'o','ArticleTools_Email_Bottom');" data-omniture="BG Gallery Share Tools - Email - Bottom" title="Share via E-mail" aria-label="Share via E-mail" href="javascript:;">
	      	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="email-blue-icon"><title id="email-blue-icon">Email to a Friend</title><g fill="#127694" fill-rule="evenodd"><path d="M21.658 4.403H2.393c2.06 1.93 8.782 8.235 9.632 9.033l9.633-9.033M15.985 11.925l6.848 6.42V5.503"/><path d="M12.026 15.638L9.24 13.025l-6.846 6.422h19.264c-1.304-1.226-4.47-4.195-6.845-6.422l-2.787 2.613zM1.22 5.503v12.842l6.846-6.42"/></g>
	<!-- fallback .png -->
	<image src="/bg-images/png/email-blue-small.png" xlink:href=""/>
</svg>
</a>
	      <div class="bg-form social-tools__email-form js-email-box">
				<h4 class="bg-form__title">Share via e-mail</h4>
				<div class="bg-form__error js-email-error"></div>
				<form name="emailForm" action="" method="post" class="js-email-form">
					<input name="story_url" type="hidden" id="story_url-Bottom" class="js-story-url" value="https://www.bostonglobe.com/news/bigpicture/2019/03/02/globe-staff-photos-month-february/zsbKlFlxkRDAiLYNucIwxJ/story.html"/>
					<label for="email-to-primary-Bottom" class="bg-form__label">To</label>
					<input name="recipient_email" type="text" id="email-to-primary-Bottom" placeholder="Separate multiple addresses with a comma" class="bg-form__input js-email-to" />
					<label for="email-mssg-primary-Bottom" class="bg-form__label">Add a message</label>
					<textarea name="message" id="email-mssg-primary-Bottom" maxlength="80" placeholder="80 character limit" class="bg-form__textarea js-email-message"></textarea>
					<label for="email-from-primary-Bottom" class="bg-form__label">Your e-mail</label>
					<input name="sender_email" type="text" id="email-from-primary-Bottom" class="bg-form__input js-email-from" />
					<input type="submit" value="Send" class="bg-button bg-button--submit js-email-submit" onClick="var s=s_gi('nytbostonglobecom');s.tl(this,'o','ArticleTools_Send_Bottom');"/>
				</form>
			</div><!-- /end .email-form -->
	    </li>
	    <li class="social-tools__list-item">
	      <a class="bg-button social-tools__link social-tools__link--facebook" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Facebook');" href="http://www.facebook.com/sharer.php?u=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25&amp;t=Globe staff photos of the month, February 2019" target="_blank" title="Share on Facebook" aria-label="Share on Facebook">
	      	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="facebook-share-icon"><title id="facebook-share-icon">Share on Facebook</title><path d="M13.503 21.63v-8.785h2.97l.445-3.425h-3.415V7.234c0-.99.277-1.667 1.71-1.667h1.826V2.503c-.317-.042-1.402-.135-2.663-.135-2.634 0-4.437 1.595-4.437 4.525V9.42H6.96v3.425h2.98v8.786h3.563" fill="#FFF" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/facebook-small.png" xlink:href=""/>
</svg>
</a>
	    </li>
	    <li class="social-tools__list-item">
	      <a class="bg-button social-tools__link social-tools__link--twitter" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Twitter');" href="http://twitter.com/intent/tweet?text=Globe staff photos of the month, February 2019&amp;url=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25&amp;via=BostonGlobe" target="_blank" title="Share on Twitter" aria-label="Share on Twitter">
	      	<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="twitter-share-icon"><title id="twitter-share-icon">Share on Twitter</title><path d="M22.457 5.614c-.77.336-1.596.563-2.464.665.886-.524 1.566-1.352 1.887-2.338-.83.484-1.747.836-2.725 1.025-.783-.82-1.898-1.334-3.132-1.334-2.37 0-4.29 1.89-4.29 4.224 0 .332.037.654.11.963-3.566-.176-6.728-1.858-8.845-4.414-.37.624-.58 1.35-.58 2.124 0 1.466.757 2.76 1.91 3.516-.705-.02-1.367-.212-1.945-.528v.053c0 2.048 1.48 3.755 3.442 4.144-.36.096-.74.148-1.13.148-.277 0-.546-.027-.808-.076.546 1.68 2.13 2.9 4.01 2.934-1.47 1.133-3.32 1.808-5.33 1.808-.347 0-.688-.02-1.024-.058 1.9 1.198 4.154 1.897 6.578 1.897 7.893 0 12.21-6.437 12.21-12.02 0-.183-.006-.365-.013-.546.838-.595 1.565-1.34 2.14-2.186" fill="#FFF" fill-rule="evenodd"/>
	<!-- fallback .png -->
	<image src="/bg-images/png/twitter-small.png" xlink:href=""/>
</svg>
</a>
	    </li class="social-tools__list-item">
	    <li class="social-tools__list-item">
			<a class="bg-button social-tools__link social-tools__link--google-plus" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Google+');" href="https://plus.google.com/share?url=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25%26s_campaign%3dsm_gp%26hl=en-US" target="_blank" title="Share on Google+" aria-label="Share on Google+">
				 		<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="google-share-icon"><title id="google-share-icon">Share on Google Plus</title><g fill="#FFF" fill-rule="evenodd"><path d="M7.82 10.955v2.508h4.15c-.17 1.076-1.255 3.156-4.15 3.156-2.497 0-4.535-2.07-4.535-4.62 0-2.55 2.038-4.62 4.535-4.62 1.42 0 2.372.607 2.916 1.13l1.985-1.912c-1.274-1.192-2.925-1.913-4.9-1.913C3.776 4.685.505 7.955.505 12c0 4.044 3.27 7.315 7.315 7.315 4.222 0 7.022-2.968 7.022-7.148 0-.48-.052-.846-.115-1.212H7.82zM23.495 10.955h-2.09v-2.09h-2.09v2.09h-2.09v2.09h2.09v2.09h2.09v-2.09h2.09"/></g>
	<!-- fallback .png -->
	<image src="/bg-images/png/google-plus-small.png" xlink:href=""/>
</svg>
</a>
				</li>
	    <li class="social-tools__list-item">
	      <a class="bg-button social-tools__link social-tools__link--linkedin" onClick="var s=s_gi('nytbostonglobecom'); s.linkTrackVars='eVar15,channel,prop1,prop4'; s.linkTrackEvents='none'; s.tl(this,'o', 'BG Article Share Tools - Linkedin');" href="http://www.linkedin.com/shareArticle?mini=true&amp;url=https%3A%2F%2Fwww.bostonglobe.com%2Fnews%2Fbigpicture%2F2019%2F03%2F02%2Fglobe-staff-photos-month-february%2FzsbKlFlxkRDAiLYNucIwxJ%2Fstory.html%3Fevent%3Devent25&amp;title=Globe staff photos of the month, February 2019&amp;summary=Here%26amp%3Brsquo%3Bs+a+look+at+some+of+the+best+images+taken+by+Globe+photographers+last+month%3A+winter+weather%2C+presidential+candidates+on+the+stump%2C+a+Chinese+New+Year+celebration%2C+Red+Sox+spring+training%2C+and+the+Patriots%26amp%3Brsquo%3B+Super+Bowl+LIII+victory+parade." target="_blank" title="Share on LinkedIn" aria-label="Share on LinkedIn"><svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-labelledby="linkedin-share-icon"><title id="linkedin-share-icon">Share on LinkedIn</title><g fill="#FFF" fill-rule="evenodd"><path d="M5.12 2.96c1.2 0 2.177.974 2.177 2.176 0 1.202-.976 2.177-2.176 2.177-1.205 0-2.177-.975-2.177-2.177S3.915 2.96 5.12 2.96zM3.24 21.04H7V8.964H3.24V21.04zM9.353 8.963h3.598v1.65h.053c.5-.948 1.726-1.95 3.55-1.95 3.802 0 4.504 2.5 4.504 5.754v6.624h-3.752v-5.873c0-1.4-.025-3.202-1.95-3.202-1.954 0-2.252 1.526-2.252 3.102v5.974h-3.75V8.964z"/></g>
	<!-- fallback .png -->
	<image src="/bg-images/png/linkedin-small.png" xlink:href=""/>
</svg>
</a>
	    </li>
	</ul>
	</aside>
<footer class="article-footer">
 <section class="article-footer">
<script>var uuid = '421735e2-3cf8-11e9-9daf-19409b78ba67';</script>
<div id="comments">
	<div id="comments-content" class="ugc-comments comments">
		<div class="loading-message" id="loading-comments">
			Loading comments...
			 </div>
	</div>
</div>
<div id="newcomments"></div>
</section>
</footer>

<div id="post-comment">
	<div id="inthisblog" class="bent">
		In this blog: Big Picture
	</div>
	<article class="entrylink">
		    <a href="/news/bigpicture/2019/03/02/globe-staff-photos-month-february/zsbKlFlxkRDAiLYNucIwxJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2019/02/05/BostonGlobe.com/Metro/Images/Chin020519PatriotsParade_Spt3.jpg">
				<div class="entrytitle">Globe staff photos of the month, February 2019</div></a>
				<div class="entrycopy geor">Boston, MA - 2/05/2019 - New England Patriots wide receiver and Super Bowl LIII MVP Julian Edelman taps his heart and points to the fans to show his love during today's New England Patriots Super Bowl LIII Victory Parade. (Barry Chin/Globe Staff), Section: Metro, Reporter: Globe Staff, Topic: 06parade, LOID: 8.5.332310335</div>
			<a href="/news/bigpicture/2019/03/02/globe-staff-photos-month-february/zsbKlFlxkRDAiLYNucIwxJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2019/02/13/smartphone-culture/GxAOC7WjTtyc8LZhs5idrN/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/09/10/BostonGlobe.com/BigPicture/Images/Rex_Byblos_International_Festival_in_9776810S.jpg">
				<div class="entrytitle">Smartphone society</div></a>
				<div class="entrycopy geor">In a world where the smartphone is now ubiquitous, photographers capture daily life in which the personal device is part of the moment. Whether it is used for communication, navigation, a flashlight, a wallet, or to take photos, it’s an essential part of life for many.</div>
			<a href="/news/bigpicture/2019/02/13/smartphone-culture/GxAOC7WjTtyc8LZhs5idrN/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2019/02/05/parade-for-champion-patriots/hB6PcfS8nkMs29dnDQbCHI/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2019/02/05/BostonGlobe.com/Metro/Images/tlumackipatsparade374.jpg">
				<div class="entrytitle">Highlights from the Patriots Super Bowl parade </div></a>
				<div class="entrycopy geor">Hundreds of thousands of Patriots fans filled the streets of Boston to celebrate the team’s sixth Super Bowl win.</div>
			<a href="/news/bigpicture/2019/02/05/parade-for-champion-patriots/hB6PcfS8nkMs29dnDQbCHI/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2019/02/03/super-bowl-liii-patriots-rams/7El5tTpVDD91ia0WqdhdGJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2019/02/04/BostonGlobe.com/BigPicture/Images/gallery lede.jpg">
				<div class="entrytitle">Super Bowl LIII: Patriots vs. Rams</div></a>
				<div class="entrycopy geor">The game was a slugfest, but New England came away with its sixth Super Bowl title.</div>
			<a href="/news/bigpicture/2019/02/03/super-bowl-liii-patriots-rams/7El5tTpVDD91ia0WqdhdGJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2019/02/01/globe-staff-photos-month-january/3XsbW2oUxOxgfMlk9Z8zON/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2019/01/24/BostonGlobe.com/Metro/Images/tlumackistormwinds322.jpg">
				<div class="entrytitle">Globe staff photos of the month, January 2019</div></a>
				<div class="entrycopy geor">&nbsp;</div>
			<a href="/news/bigpicture/2019/02/01/globe-staff-photos-month-january/3XsbW2oUxOxgfMlk9Z8zON/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2019/01/15/years-ago-boston-north-end-was-hit-deadly-wave-molasses/xyn8sjJG3gYt3ZcixyuaYO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2019/01/07/BostonGlobe.com/BigPicture/Images/150114_MJ_002.jpg">
				<div class="entrytitle">100 years ago, Boston’s North End was hit by a deadly wave of molasses</div></a>
				<div class="entrycopy geor">Millions of gallons of brown syrup leveled buildings and buckled the steel girders of the elevated railway, drowning unsuspecting people in its wake.</div>
			<a href="/news/bigpicture/2019/01/15/years-ago-boston-north-end-was-hit-deadly-wave-molasses/xyn8sjJG3gYt3ZcixyuaYO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2019/01/04/the-best-boston-globe-staff-photos/6lvDHDoDGqTl2M8oB9QfTN/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/12/07/BostonGlobe.com/Magazine/Images/RinaldiFeature02-004.jpg">
				<div class="entrytitle">The best Boston Globe staff photos of 2018 </div></a>
				<div class="entrycopy geor">&nbsp;</div>
			<a href="/news/bigpicture/2019/01/04/the-best-boston-globe-staff-photos/6lvDHDoDGqTl2M8oB9QfTN/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2019/01/03/globe-staff-photos-month-december/sOPbjqicOMFy9BwHqHUTBI/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/12/11/BostonGlobe.com/ReceivedContent/Images/KREITER12112018ChildrensHoliday1.jpg">
				<div class="entrytitle">Globe staff photos of the month, December 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: celebrating the holidays, honoring a fallen firefighter, the death of President George H.W. Bush, remembering Pearl Harbor, and high school football championships.</div>
			<a href="/news/bigpicture/2019/01/03/globe-staff-photos-month-december/sOPbjqicOMFy9BwHqHUTBI/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/12/29/double-diagnosis/IUVSDlylfEIuCczz4Il6uJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/11/30/BostonGlobe.com/Metro/Images/Walker_101118_xxpoorwithcancer_4286xx.jpg">
				<div class="entrytitle">A double diagnosis — cancer while poor</div></a>
				<div class="entrycopy geor">&nbsp;</div>
			<a href="/news/bigpicture/2018/12/29/double-diagnosis/IUVSDlylfEIuCczz4Il6uJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/12/26/the-year-pictures-part/GdY9DGnrm4KROzkuH9IZxH/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/12/18/BostonGlobe.com/BigPicture/Images/AFP_19K5I5.jpg">
				<div class="entrytitle">The year 2018 in pictures: Part II</div></a>
				<div class="entrycopy geor">Photographs from July to December on a range of topics from around the world.</div>
			<a href="/news/bigpicture/2018/12/26/the-year-pictures-part/GdY9DGnrm4KROzkuH9IZxH/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/12/20/the-year-pictures-part/oNSsXtYoFphoTeQetCsgNP/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/12/18/BostonGlobe.com/BigPicture/Images/973077552-7519.jpg">
				<div class="entrytitle">The year 2018 in pictures: Part I</div></a>
				<div class="entrycopy geor">Photographs from January to June on a range of topics from around the world.</div>
			<a href="/news/bigpicture/2018/12/20/the-year-pictures-part/oNSsXtYoFphoTeQetCsgNP/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/12/04/globe-staff-photos-month-november/5PnoNz67cCEuyq9Z72ej9H/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/11/30/BostonGlobe.com/Metro/Images/ryanwintertreemet.jpg">
				<div class="entrytitle">Globe staff photos of the month, November 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: a colorful sunrise, the start of recreational marijuana sales, Christmas tree lightings, Alex Cora visits Puerto Rico, and high school sports.</div>
			<a href="/news/bigpicture/2018/12/04/globe-staff-photos-month-november/5PnoNz67cCEuyq9Z72ej9H/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/11/30/keeping-tradition-hunt-alive/08jbDctL0QB1o2Bq6ma3iJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/11/26/BostonGlobe.com/Lifestyle/Images/RinaldiFoxHunt108.jpg">
				<div class="entrytitle">Keeping tradition of the hunt alive</div></a>
				<div class="entrycopy geor">Through the New England countryside, riders take part in an adventure at once ancient and modern: A formal fox hunt, no fox necessary.</div>
			<a href="/news/bigpicture/2018/11/30/keeping-tradition-hunt-alive/08jbDctL0QB1o2Bq6ma3iJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/11/26/chaos-border-mexico/m1Ze1NTyajTqBhwLZy2LvK/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/11/26/BostonGlobe.com/BigPicture/Images/AFP_1B39U5-7412.jpg">
				<div class="entrytitle">Chaos at the US border in Mexico</div></a>
				<div class="entrycopy geor">A peaceful march devolved into unrest when US agents fired tear gas to stop hundreds of migrants attempting to storm a border fence separating Mexico from the United States. Over 5,000 Central American migrants have been camping out at a sports complex in Tijuana.</div>
			<a href="/news/bigpicture/2018/11/26/chaos-border-mexico/m1Ze1NTyajTqBhwLZy2LvK/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/11/15/the-deadly-camp-fire-california/E8rvsaRtrYxDmnswjXV8RO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/11/15/BostonGlobe.com/BigPicture/Images/AFP_1AV45J-7353.jpg">
				<div class="entrytitle">The deadly Camp Fire in California</div></a>
				<div class="entrycopy geor">At least 63 people have lost their lives in California’s deadliest wildfire. Over 10,000 structures were destroyed. The search for hundreds of people still missing continues.</div>
			<a href="/news/bigpicture/2018/11/15/the-deadly-camp-fire-california/E8rvsaRtrYxDmnswjXV8RO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/11/09/anniversary-end-wwi/xcv5eE7gIpSVNFmJfTRCpL/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/11/09/BostonGlobe.com/BigPicture/Images/54c05ecca4454a60b729c40af07ac488-54c05ecca4454a60b729c40af07ac488-0.jpg">
				<div class="entrytitle">100th anniversary of the end of WWI</div></a>
				<div class="entrycopy geor">In this Nov. 1918 file photo, American soldiers from New York, who served on the frontline in Cambria, France, rig up a Liberty Bell to celebrate the signing of the Armistice to end World War One.</div>
			<a href="/news/bigpicture/2018/11/09/anniversary-end-wwi/xcv5eE7gIpSVNFmJfTRCpL/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/11/04/globe-staff-photos-month-october/PrrfKGWnKkOnDFWyxlQInO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/11/BostonGlobe.com/Metro/Images/ryan_fallcolors2_met.jpg">
				<div class="entrytitle">Globe staff photos of the month, October 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: fall colors, Halloween, the Milken Educator Award winner, Harvard University lawsuit, and a World Series victory.</div>
			<a href="/news/bigpicture/2018/11/04/globe-staff-photos-month-october/PrrfKGWnKkOnDFWyxlQInO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/31/red-sox-world-series-victory-parade/YIE9XJPdrcS2xWbQaIF3aJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/31/BostonGlobe.com/Sports/Images/tlumackiSOXparade194.jpg">
				<div class="entrytitle">Red Sox 2018 World Series victory parade</div></a>
				<div class="entrycopy geor">The Red Sox once again for the fourth time this century rolled through the city in triumph after winning the 2018 World Series.</div>
			<a href="/news/bigpicture/2018/10/31/red-sox-world-series-victory-parade/YIE9XJPdrcS2xWbQaIF3aJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/28/world-series-game-red-sox-dodgers/0QHQvfeDAVz5c7LGNGdbnI/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/29/BostonGlobe.com/ReceivedContent/Images/WS5stan22.jpg">
				<div class="entrytitle">Red Sox win the 2018 World Series</div></a>
				<div class="entrycopy geor">The Boston Red Sox took on the Los Angeles Dodgers in California for the 2018 World Series. The Red Sox defeated the Dodgers 5-1, to capture their ninth World Series title.</div>
			<a href="/news/bigpicture/2018/10/28/world-series-game-red-sox-dodgers/0QHQvfeDAVz5c7LGNGdbnI/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/27/world-series-game-red-sox-dodgers/BWzkYCd9yz9m2Q75zBjlqJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/28/BostonGlobe.com/ReceivedContent/Images/WS4davis24.jpg">
				<div class="entrytitle">World Series: Game 4, Red Sox vs. Dodgers</div></a>
				<div class="entrycopy geor">The Boston Red Sox take on the Los Angeles Dodgers in California during the 2018 World Series. The Sox won Game 4, 9-6.</div>
			<a href="/news/bigpicture/2018/10/27/world-series-game-red-sox-dodgers/BWzkYCd9yz9m2Q75zBjlqJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/26/world-series-game-red-sox-dodgers/PRP3ZmSYmIHUHsJMwuIA9K/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/27/BostonGlobe.com/ReceivedContent/Images/102718Grossfeld54-001.jpg">
				<div class="entrytitle">World Series: Game 3, Red Sox vs. Dodgers</div></a>
				<div class="entrycopy geor">The Boston Red Sox take on the Los Angeles Dodgers in California during the 2018 World Series.</div>
			<a href="/news/bigpicture/2018/10/26/world-series-game-red-sox-dodgers/PRP3ZmSYmIHUHsJMwuIA9K/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/24/world-series-game-red-sox-dodgers/fvW7KTRdHEPm0bu1FD38uO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/25/BostonGlobe.com/ReceivedContent/Images/WS2davis18.jpg">
				<div class="entrytitle">World Series: Game 2, Red Sox vs. Dodgers</div></a>
				<div class="entrycopy geor">The Boston Red Sox take on the Los Angeles Dodgers at Fenway Park in the 2018 World Series. Red Sox win 4-2 and lead series 2-0.</div>
			<a href="/news/bigpicture/2018/10/24/world-series-game-red-sox-dodgers/fvW7KTRdHEPm0bu1FD38uO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/23/world-series-game-sox-dodgers/N8VDneWCm83SCSM6BdpoTO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/24/BostonGlobe.com/ReceivedContent/Images/102318Tlumacki1.jpg">
				<div class="entrytitle">World Series: Game 1, Red Sox vs. Dodgers</div></a>
				<div class="entrycopy geor">The Boston Red Sox take on the Los Angeles Dodgers in the 2018 World Series.</div>
			<a href="/news/bigpicture/2018/10/23/world-series-game-sox-dodgers/N8VDneWCm83SCSM6BdpoTO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/03/globe-staff-photos-month-september/7me0Qfv9hRUwC9DbqsvnBN/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/10/02/BostonGlobe.com/BigPicture/Images/T6FJX856.jpg">
				<div class="entrytitle">Globe staff photos of the month, September 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: gas explosions in Lawrence, 9/11 remembrance ceremonies, primary elections, somber pets, the New England Giant Pumpkin Weigh-Off at the Topsfield Fair, and the beginning of the Patriots regular season.</div>
			<a href="/news/bigpicture/2018/10/03/globe-staff-photos-month-september/7me0Qfv9hRUwC9DbqsvnBN/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/10/02/scenes-from-aftermath-indonesia-earthquake-tsunami/FraOkK9sxfPCqM4TyZ5WPJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/08/08/BostonGlobe.com/National/Images/5b1caec251004aeda737b4ed2a4ab53a-5b1caec251004aeda737b4ed2a4ab53a-0.jpg">
				<div class="entrytitle">Scenes from the aftermath of Indonesia’s earthquake, tsunami</div></a>
				<div class="entrycopy geor">&nbsp;</div>
			<a href="/news/bigpicture/2018/10/02/scenes-from-aftermath-indonesia-earthquake-tsunami/FraOkK9sxfPCqM4TyZ5WPJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/09/22/the-world-stage-way-ahead/KjkIVw6IXOJo31civWdzmO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/08/02/BostonGlobe.com/Metro/Images/RinaldiDrama01.jpg">
				<div class="entrytitle">The world, the stage, the way ahead</div></a>
				<div class="entrycopy geor">Deb was still new to this high school in Peabody. She had arrived a year ago, a stranger dropped into the junior class, knowing no one and lacking all their shared history. She had found her way to the vaunted theater program, establishing a foothold in a realm that felt magical. This show was headed to a high-stakes statewide competition. Did she really think she belonged on that stage? She had resolved to find out.</div>
			<a href="/news/bigpicture/2018/09/22/the-world-stage-way-ahead/KjkIVw6IXOJo31civWdzmO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/09/16/pipeline-from-africa/Vgi8t2v77qhHivGsay7N5N/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/08/07/BostonGlobe.com/Metro/Images/Globe_arrival.jpg">
				<div class="entrytitle">Pipeline from Africa</div></a>
				<div class="entrycopy geor">Immigrants from Ghana do much of the low-paying, back-breaking work of caring for frail Americans in their homes. Back home, they’re seen as success stories.</div>
			<a href="/news/bigpicture/2018/09/16/pipeline-from-africa/Vgi8t2v77qhHivGsay7N5N/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/09/13/hurricane-florence/Faby8PfsCOcSWvgaqptaSM/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/09/17/BostonGlobe.com/BigPicture/Images/1033074056-7216.jpg">
				<div class="entrytitle">The destructive aftermath of Hurricane Florence </div></a>
				<div class="entrycopy geor">Hurricane Florence lashed the Southeast coast of the United States. The storm dumped heavy amounts of rain across the region and killed at least 32 people.</div>
			<a href="/news/bigpicture/2018/09/13/hurricane-florence/Faby8PfsCOcSWvgaqptaSM/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/09/05/globe-staff-photos-month-august/dcBNtaIUJYe4KFaiEB5HcP/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/08/07/BostonGlobe.com/Metro/Images/Walker_080718_08heat_3832x.jpg">
				<div class="entrytitle">Globe staff photos of the month, August 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: staying cool in the summer heat, installing a new police commissioner, spotting sharks on the Cape, campaigning politicians on the trail, and flipping gymnasts at the US championships.</div>
			<a href="/news/bigpicture/2018/09/05/globe-staff-photos-month-august/dcBNtaIUJYe4KFaiEB5HcP/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/08/25/home-but-for-how-long/8z3bNGjDqksbe3zbqaVQtK/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/08/15/BostonGlobe.com/Metro/Images/RinaldiPatricia49.jpg">
				<div class="entrytitle">A home, but for how long?</div></a>
				<div class="entrycopy geor">For 19 years, the US government had given Patricia Carbajal permission to stay in this country, to work, to put down roots. For 19 years, administration after administration extended Temporary Protected Status for Honduras after the destruction wrought by Hurricane Mitch in 1998 was compounded by crippling poverty, destabilizing corruption, and violence so pervasive that the murder rate in Honduras is now among the highest in the world. After 19 years, Patricia’s status had long ago stopped feeling temporary.  But, now, in a moment, everything could change.</div>
			<a href="/news/bigpicture/2018/08/25/home-but-for-how-long/8z3bNGjDqksbe3zbqaVQtK/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/08/16/gymnastics-championships-boston/0QrdRRtbpsMSH2gapWvsFP/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/08/16/BostonGlobe.com/ReceivedContent/Images/g13.jpg">
				<div class="entrytitle">US Gymnastics Championships in Boston</div></a>
				<div class="entrycopy geor">This week, Boston will get a look at some of the best gymnasts in the world during the US Nationals, which runs Thursday through Sunday at TD Garden.</div>
			<a href="/news/bigpicture/2018/08/16/gymnastics-championships-boston/0QrdRRtbpsMSH2gapWvsFP/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/08/07/summertime-festivals-around-world/vmkz0rfqEUTJVynpJD6ihP/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/08/03/BostonGlobe.com/BigPicture/Images/AFP_1813QE.jpg">
				<div class="entrytitle">Summertime festivals around the world</div></a>
				<div class="entrycopy geor">Photographers capture communities gathering at events organized to celebrate a variety of cultural observances.</div>
			<a href="/news/bigpicture/2018/08/07/summertime-festivals-around-world/vmkz0rfqEUTJVynpJD6ihP/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/08/01/globe-staff-photos-month-july/XDxpj4ayHxxw1QdBTkSgFK/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/07/22/BostonGlobe.com/Arts/Images/lee_072118_grohl5_liv.jpg">
				<div class="entrytitle">Globe staff photos of the month, July 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: the Puerto Rican Parade, Taylor Swift performing at Gillette Stadium, celebrating fourth of July, mourning a fallen police officer, loggerhead turtle release, and Patriots training camp</div>
			<a href="/news/bigpicture/2018/08/01/globe-staff-photos-month-july/XDxpj4ayHxxw1QdBTkSgFK/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/07/26/the-life-aquatic/OOKXBNaHGKpKpdXpMqYAFI/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/07/26/BostonGlobe.com/BigPicture/Images/Rex_Marine_life_at_the_Two_Oceans_Aq_9767593Q-7046.jpg">
				<div class="entrytitle">The life aquatic</div></a>
				<div class="entrycopy geor">During the summer heat, a look at animals cooling off with water or in their aquatic environments.</div>
			<a href="/news/bigpicture/2018/07/26/the-life-aquatic/OOKXBNaHGKpKpdXpMqYAFI/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/07/10/all-eyes-world-cup/mUjbtDb9MOKGlx3OiPPrpL/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/06/22/BostonGlobe.com/BigPicture/Images/AFP_15Z148.jpg">
				<div class="entrytitle">All eyes on the World Cup</div></a>
				<div class="entrycopy geor">Around the world, billions of fans are watching the 2018 FIFA World Cup, which is being held in Russia for the first time. The tournament features thirty-two teams from six continents.</div>
			<a href="/news/bigpicture/2018/07/10/all-eyes-world-cup/mUjbtDb9MOKGlx3OiPPrpL/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/07/07/kakuma-refugee-camp/WA4vybjw2gpDMns0bn5VpN/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/06/11/BostonGlobe.com/ReceivedContent/Images/BEDFORD_20180511_KAKUMA_028.jpg">
				<div class="entrytitle">Kakuma refugee camp</div></a>
				<div class="entrycopy geor">&nbsp;</div>
			<a href="/news/bigpicture/2018/07/07/kakuma-refugee-camp/WA4vybjw2gpDMns0bn5VpN/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/07/02/globe-staff-photos-month-june/gNw9mgCYPuZkQ0YfKZ18YN/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/06/20/BostonGlobe.com/Sports/Images/Chin062018D4StateBaseballSemifinalArchbishopWilliamsHamiltonWenham_Spt1.jpg">
				<div class="entrytitle">Globe staff photos of the month, June 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: watching World Cup soccer, The annual Boston Pride Parade, protesting gun volence, protesting the presidents immigration border policy, blessing animals, ad great action - reaction on local baseball diamonds..</div>
			<a href="/news/bigpicture/2018/07/02/globe-staff-photos-month-june/gNw9mgCYPuZkQ0YfKZ18YN/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/06/25/life-bleachers-fenway-park/jz3Qwm5h7dhymYbXa4l7ZO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/06/01/BostonGlobe.com/ReceivedContent/Images/bleacher563.jpg">
				<div class="entrytitle">Life in the bleachers at Fenway Park</div></a>
				<div class="entrycopy geor">Fans climbed the stairs to the upper bleachers at Fenway Park.</div>
			<a href="/news/bigpicture/2018/06/25/life-bleachers-fenway-park/jz3Qwm5h7dhymYbXa4l7ZO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/06/19/border-policy-controversy/co40wcQZdejrWDb5tIoEEK/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/06/18/BostonGlobe.com/BigPicture/Images/973077450.jpg">
				<div class="entrytitle">US border policy controversy</div></a>
				<div class="entrycopy geor">Customs and border patrol officials in the United States are enforcing President Trump’s strict immigration policies on the Mexico border, causing widespread protests. At least 2,000 children have been separated from their parents since April.</div>
			<a href="/news/bigpicture/2018/06/19/border-policy-controversy/co40wcQZdejrWDb5tIoEEK/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/06/08/opioids-land-more-women-behind-bars/XK4hlaJ8tCxrYD1Umj0t1I/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/06/08/BostonGlobe.com/BigPicture/Images/b2e8c615361f4f6b950f23a5ef7402af-8f4c9dd1bfa34482b7548f55c73cce25-0.jpg">
				<div class="entrytitle">Opioids land more women behind bars</div></a>
				<div class="entrycopy geor">This lone county jail in a remote corner of Appalachia offers an agonizing glimpse into how the tidal wave of opioids and methamphetamines has ravaged America.</div>
			<a href="/news/bigpicture/2018/06/08/opioids-land-more-women-behind-bars/XK4hlaJ8tCxrYD1Umj0t1I/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/06/01/globe-staff-photos-month-may/hTXn3RrtjmIGCeVBXkFjPM/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/05/09/BostonGlobe.com/Metro/Images/tlumacki_acresandacresoftulips_metro555.jpg">
				<div class="entrytitle">Globe staff photos of the month, May 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: a tulip farm in Rhode Island, local graduations, Memorial Day, Boston Calling Music Festival, and the NHL and NBA playoffs.</div>
			<a href="/news/bigpicture/2018/06/01/globe-staff-photos-month-may/hTXn3RrtjmIGCeVBXkFjPM/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/05/21/eruption-hawaii-kilauea-volcano/rBCkBHoRGTM8l3RaWUITfJ/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/05/21/BostonGlobe.com/BigPicture/Images/Rex_Massive_fast_move_lava_flow_from_9685940B.jpg">
				<div class="entrytitle">Eruption of Hawaii’s Kilauea volcano </div></a>
				<div class="entrycopy geor">Bruce Omori/Paradise Helicopters/EPA/Shutterstock</div>
			<a href="/news/bigpicture/2018/05/21/eruption-hawaii-kilauea-volcano/rBCkBHoRGTM8l3RaWUITfJ/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/05/04/raising-connor/ssKpjPR8YrmYgHWzDYl3cO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/03/28/BostonGlobe.com/ReceivedContent/Images/Walker_090317_menatlhealth_6746xA.jpg">
				<div class="entrytitle">Raising Connor</div></a>
				<div class="entrycopy geor">He is easy to love, affectionate, and friendly. He is moody and unpredictable. Vulnerable, sweet, devoted to family. Impulsive, strong, and overflowing with emotion. Dreaming of home, always. Never quite at home, anywhere. This is Connor.</div>
			<a href="/news/bigpicture/2018/05/04/raising-connor/ssKpjPR8YrmYgHWzDYl3cO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/05/02/globe-staff-photos-month-april/ogdHo3c9z4HMT2DtXNdP3L/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/04/05/BostonGlobe.com/Sports/Images/Walker_040518_RedSox_Rays_7490x.jpg">
				<div class="entrytitle">Globe staff photos of the month, April 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: a funeral for a fallen police officer, the Boston Marathon, a new home for a lost dog, opening day at Fenway Park, and the Bruins and Celtics in the playoffs.</div>
			<a href="/news/bigpicture/2018/05/02/globe-staff-photos-month-april/ogdHo3c9z4HMT2DtXNdP3L/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/04/20/spring-blossoms/V7sAbKC5XvQK74oZaqojRO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/04/20/BostonGlobe.com/BigPicture/Images/941344788-6634.jpg">
				<div class="entrytitle">Spring blossoms</div></a>
				<div class="entrycopy geor">To commemorate Earth Day on April 22, a look at transforming landscapes around the world bursting with color as warm weather approaches.</div>
			<a href="/news/bigpicture/2018/04/20/spring-blossoms/V7sAbKC5XvQK74oZaqojRO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/04/16/boston-marathon/nsoWo6i6bFjM77Xwis5PRI/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/04/16/BostonGlobe.com/Sports/Images/tlumacki_bostonmarathonfinishline_sports1405.jpg">
				<div class="entrytitle">2018 Boston Marathon</div></a>
				<div class="entrycopy geor">Thousands of runners compete under miserable weather conditions during the 122nd running of the Boston Marathon.</div>
			<a href="/news/bigpicture/2018/04/16/boston-marathon/nsoWo6i6bFjM77Xwis5PRI/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/04/03/remembering-martin-luther-king-april/9GcOZCbFbNlR5fStF13wDK/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/04/03/BostonGlobe.com/ReceivedContent/Images/AP_6804091863-001.jpg">
				<div class="entrytitle">Remembering Martin Luther King Jr.</div></a>
				<div class="entrycopy geor">A look at some of the pivotal moments in the life of civil rights leader Martin Luther King Jr. as we mark the 50th anniversary of the tragic end of his life on April 4, 1968.</div>
			<a href="/news/bigpicture/2018/04/03/remembering-martin-luther-king-april/9GcOZCbFbNlR5fStF13wDK/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/04/02/globe-staff-photos-month-march/fx1OqcVbKEyCnPCqqMePfK/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/03/02/BostonGlobe.com/Metro/Images/tlumacki_stormflooding_metro708.jpg">
				<div class="entrytitle">Globe staff photos of the month, March 2018</div></a>
				<div class="entrycopy geor">Here’s a look at some of the best images taken by Globe photographers last month: multiple nor’easters, St. Patrick’s Day, protesting gun violence, high school state championships for hockey and basketball, Good Friday, and the start of Red Sox season</div>
			<a href="/news/bigpicture/2018/04/02/globe-staff-photos-month-march/fx1OqcVbKEyCnPCqqMePfK/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/03/16/paralympic-winter-games/hHVXJtIdly8xVLTDQLC4EO/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/03/16/BostonGlobe.com/BigPicture/Images/2f5c7201abc54e6baf090e04dd1883e6-2f5c7201abc54e6baf090e04dd1883e6-0.jpg">
				<div class="entrytitle">2018 Paralympic Winter Games</div></a>
				<div class="entrycopy geor">Scenes from the Paralympics taking place March 9-18 in PyeongChang, South Korea. 670 athletes with disabilities from around the world compete in 80 events in six different sports.</div>
			<a href="/news/bigpicture/2018/03/16/paralympic-winter-games/hHVXJtIdly8xVLTDQLC4EO/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		<article class="entrylink">
		    <a href="/news/bigpicture/2018/03/08/back-back-nor-easter-storms-slam-mass/OP4gvvXiFZ4n92WgPrCkeK/story.html?p1=Gallery_InThisSection_Bottom" onclick="trackBlog();">
				<img src="//c.o0bg.com/rf/image_90x90/Boston/2011-2020/2018/03/08/BostonGlobe.com/Metro/Images/tlumacki_stormflooding_baby.jpg">
				<div class="entrytitle">Back to back nor’easter storms slam Massachusetts</div></a>
				<div class="entrycopy geor">Two nor’easter storms in the past week have brought snow, power outages and flooding to towns across Massachusetts.</div>
			<a href="/news/bigpicture/2018/03/08/back-back-nor-easter-storms-slam-mass/OP4gvvXiFZ4n92WgPrCkeK/story.html?p1=Gallery_InThisSection_Bottom" class="goto helv">GO TO ENTRY</a>
		</article>
		</div>
<!-- Page footer -->
<div class="subscription-banner">
	<a class="subscription-banner__link" href="http://subscribe.bostonglobe.com/B6704/?p1=BGFooter_DigitalSubscription_Banner">
		<span class="subscription-banner__title">Real journalists. Real journalism. </span>
		<span class="subscription-banner__cta" >Subscribe to The Boston Globe today.</span>
	</a>
</div>
<footer class="site-footer" role="contentinfo">
	<div class="site-footer__content">
		<ul class="site-footer__list">
			<li class="site-footer__list-item"><h3 class="site-footer__heading">Subscribe Now</h3></li>
			<li class="site-footer__list-item"><a class="site-footer__link" href="https://subscribe.bostonglobe.com/B5759/?p1=BGFooter">Digital Access</a></li>
			<li class="site-footer__list-item"><a class="site-footer__link" href="https://subscribe.bostonglobe.com/B3428/?p1=BGFooter">Home Delivery</a></li>
			<li class="site-footer__list-item"><a class="site-footer__link" href="https://manage.bostonglobe.com/GiftTheGlobe/LandingPage.html">Gift Subscriptions</a></li>
			<li class="site-footer__list-item"><h3 class="site-footer__heading">My Account</h3></li>
			<li class="site-footer__list-item"><a href="https://manage.bostonglobe.com/cs/mc/login.aspx?p1=BGFooter" class="site-footer__link">Manage my Account</a></li>
			<li class="site-footer__list-item"><a href="https://www.bostonglobe.com/bgcs" class="site-footer__link" onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Footer - Download Customer Service App');">Mobile Customer Service</a></li>
			<li class="site-footer__list-item"><a href="/newsletters?p1=BGFooter_Newsletters" class="site-footer__link">Sign Up For Newsletters</a></li>
		</ul>

		<ul class="site-footer__list">
			<li class="site-footer__list-item"><h3 class="site-footer__heading">Contact</h3></li>
			<li class="site-footer__list-item"><a href="https://customerservice.globe.com/hc/?p1=BGFooter" class="site-footer__link">Help</a></li>
			<li class="site-footer__list-item"><a href="https://customerservice.globe.com/hc/?p1=BGFooter" class="site-footer__link">FAQs</a></li>
			<li class="site-footer__list-item"><a href="/tools/help/stafflist?p1=BGFooter" class="site-footer__link">Globe newsroom</a></li>
			<li class="site-footer__list-item"><a href="https://www.bostonglobemedia.com/" class="site-footer__link" onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Footer - BGM Home Page');">Advertise</a></li>
			<li class="site-footer__list-item"><a href="https://manage.bostonglobe.com/Order/newspaper/Newspaper.aspx" class="site-footer__link" onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Footer - Order Back Issues');">Order back issues</a></li>
		</ul>

		<ul class="site-footer__list">
			<li class="site-footer__list-item"><h3 class="site-footer__heading">Social</h3></li>
			<li class="site-footer__list-item"><a href="https://www.facebook.com/globe" class="site-footer__link">Facebook</a></li>
			<li class="site-footer__list-item"><a href="https://twitter.com/#!/BostonGlobe" class="site-footer__link">Twitter</a></li>
			<li class="site-footer__list-item"><a href="https://plus.google.com/108227564341535363126/about" class="site-footer__link">Google+</a></li>
		</ul>

		<ul class="site-footer__list">
			<li class="site-footer__list-item"><h3 class="site-footer__heading">More</h3></li>
		    <li class="site-footer__list-item"><a href="https://epaper.bostonglobe.com/launch.aspx?pbid=2c60291d-c20c-4780-9829-b3d9a12687cf" class="site-footer__link" onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Footer - EPaper launch');">ePaper</a></li>
			<li class="site-footer__list-item"><a href="http://nieonline.com/bostonglobe/" class="site-footer__link" onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Footer - News In Education');">News in Education</a></li>
			<li class="site-footer__list-item"><a href="https://www.bostonglobe.com/eom/SysConfig/WebPortal/BostonGlobe/Framework/regi/archive-login.jsp" class="site-footer__link" onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Footer - Public Archives');">Archives</a></li>
			<li class="site-footer__list-item"><a href="/tools/help/privacy?p1=BGFooter" class="site-footer__link">Privacy policy</a></li>
			<li class="site-footer__list-item"><a href="/tools/help/terms-service?p1=BGFooter" class="site-footer__link">Terms of service</a></li>
			<li class="site-footer__list-item"><a href="/termsofpurchase?p1=BGFooter" class="site-footer__link">Terms of purchase</a></li>
			<li class="site-footer__list-item"><a href="https://www.bostonglobemedia.com/careers" class="site-footer__link" onclick="var s=s_gi('nytbostonglobecom');s.linkTrackVars='channel,prop1,eVar15';s.linkTrackEvents='none';s.tl(this,'o','BG Footer - BGM Careers');">Work at Boston Globe Media</a></li>
		</ul>
		<p class="site-footer__copyright">&copy; 2019 Boston Globe Media Partners, LLC</p>
	</div>
</footer>

<script type="text/javascript">
	var google_conversion_id = 1071256246;
	var google_custom_params = window.google_tag_params;
	var google_remarketing_only = true;
</script>
<script type="text/javascript" src="//www.googleadservices.com/pagead/conversion.js" ></script>
<noscript>
	<div style="display:inline;">
		<img height="1" width="1" style="border-style:none;" alt="" src="//googleads.g.doubleclick.net/pagead/viewthroughconversion/1071256246/?value=0&amp;guid=ON&amp;script=0"/>
	</div>
</noscript>
<div class="ad ad-slot-c" id="EXTRA"></div><!-- Tracking -->
<!-- stubView3 =  -->
<!-- sectionpath /news/bigpicture -->



<!-- pageTitle:  -->

<!-- Determine sprop_3 -->




	
	

<script>
	var sPageName = (typeof sPageNameOverride != 'undefined') ? sPageNameOverride : 'News | The Big Picture | Globe staff photos of the month, February 2019';
	//if (typeof sPageNameOverride != 'undefined') {
	//	sPageName = sPageNameOverride;
	//}
	var sChannel = (typeof sChannelOverride != 'undefined') ? sChannelOverride : 'News';
	//var sChannel  = 'News';
	var sProp1    = 'News | The Big Picture';
	var sProp3    = '';
	var sProp4    = 'Standard';
	//var sProp6    = 'Photo Roll | The Big Picture';
	var sProp6    = (typeof sProp6Val != 'undefined') ? sProp6Val : 'Photo Roll | The Big Picture';
	var sProp8	  = '03/02/2019 02:33:56 PM';
	var sProp16    = '';
	var sProp67	  = '421735e2-3cf8-11e9-9daf-19409b78ba67';
	var sProp68	  = 'Globe staff photos of the month, February 2019';
	var sProp6a = sProp6;
	var sProp1a = sProp1;
	var sPageNamea = sPageName;
	var taxonomies = '';
	var sChannela = sChannel;
	  if ( methode.showPaywall ) {
	    sChannel  = 'Member Center';
	    sProp1    = 'Member Center | BGC Registration';
	    if ( $(window).width() <= 768 ) {
	      // assume mobile
	      sPageName = 'Member Center | BGC Registration | Fullpage Paywall Challenge';
	      sProp6    = 'BGC Registration Page - Fullpage';
	    } else {
	      // assume desktop
	      sPageName = 'Member Center | BGC Registration | Modal Paywall Challenge';
	      sProp6    = 'BGC Registration Page - Modal';
	    }
	  }
	globe.extend(globe.analytics.omniture, {
	  'pageType' : '',
	  'pageName' : sPageName,
	  'channel'  : sChannel,
	  'prop1'    : sProp1,
	  'prop3'    : sProp3,
	  'prop4'    : sProp4,
	  'prop6'    : sProp6,
	  'prop8'    : sProp8,
	  'prop16'    : sProp16,
	  
		'prop35' : 'logged out',
		'eVar20' : 'logged out',
   		'prop42' : 'Metered',
   		'eVar42' : 'Metered',

	  'eVar16'   : '',
	  'eVar41'   : 'BostonGlobe.com',
	  'eVar41'   : 'BostonGlobe.com',
	  'eVar67'   : sProp67,
	  'prop67'    : sProp67,
	  'prop68'    : sProp68,
	  'list2'	 : taxonomies,
	  'WT omnituretracking.tag' : 'big-picture'
	});
</script>

<noscript>
	<img src="https://newyorktimes.112.2o7.net/b/ss/nytbostonglobecom/1/H.23.3--NS/0?pageName=News%20%7C%20BGC%20Homepage&amp;ch=News&amp;c1=News%20%7C%20BGC%20Homepage&amp;c6=BGC%20Homepage&amp;c41=BostonGlobe.com&amp;v41=BostonGlobe.com" height="1" width="1" border="0" alt="" />
</noscript>


</div> <!-- end #contain.content -->

<script src="/js/dist/common.js,globe-bgbrandlab-belt.min.js,dist/globe-analytics.js,globe-arcvideo-preroll.js?v=19745I2111"></script>
<link rel="stylesheet" type="text/css" media="print" href="/css/globe-print.css?v=19745I2111"/>
<link rel="stylesheet" type="text/css" href="//meter.bostonglobe.com/css/style.css">
<script src="//meter.bostonglobe.com/js/meter.js"></script>
<script src="/js/dist/SiteHeaderButtons.js,dist/UserStateChange.js,dist/globe-blueconic.js,dist/chartbeat.js,dist/globe-subscription-banner.js,dist/globe-adinclude.js?v=19745I2111"></script>
<link rel="stylesheet" type="text/css" media="print" href="/css/globe-print.css?v=19745I2111"/>
<script type="text/javascript">
	(function() {
		if (typeof window.meterObject !== 'undefined') {
			
			// Change the meter base URL (this allows all campaign links to be relative) for this particular environment.
			window.meterObject.updateBaseURL('https://www.bostonglobe.com');
			
			window.meterObject.executeMeter({
				loid:'5.0.1114829766',
				uuid:'421735e2-3cf8-11e9-9daf-19409b78ba67',
				
				exemptFromMeter: false
			});
			
		}

	})();
</script>
</body>
</html>
`

func TestBigPictureFeedBuiler(t *testing.T) {
	node, _ := goquery.NewDocumentFromReader(strings.NewReader(file))
	handler := BigPictureFeed{}

	resp, err := handler.ParseFeed(node)
	if err != nil {
		t.Error("Should parse the file")
	}
	if resp.Title != "Globe staff photos of the month, February 2019" {
		t.Error("Titles didn't match. Found: %s", resp.Title)
	}
}
