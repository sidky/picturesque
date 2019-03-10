package feed

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

const infocusFile = `









<!doctype html>

<html class="no-js" lang="en" prefix="og: http://ogp.me/ns#">

<head data-template-set="html5-reset" prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# article: http://ogp.me/ns/article#">

    <meta charset="utf-8">

    
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set - The Atlantic</title>

    

        
        <link rel="shortcut icon" href="https://cdn.theatlantic.com/assets/static/a/theatlantic/common/img/favicon.ico">
        <link rel="apple-touch-icon" href="https://cdn.theatlantic.com/assets/static/a/theatlantic/common/img/apple-touch-icon-iphone.png">
        <link rel="apple-touch-icon" sizes="76x76" href="https://cdn.theatlantic.com/assets/static/a/theatlantic/common/img/apple-touch-icon-ipad.png">
        <link rel="apple-touch-icon" sizes="120x120" href="https://cdn.theatlantic.com/assets/static/a/theatlantic/common/img/apple-touch-icon-iphone-retina.png">
        <link rel="apple-touch-icon" sizes="152x152" href="https://cdn.theatlantic.com/assets/static/a/theatlantic/common/img/apple-touch-icon-ipad-retina.png">
        <link rel="mask-icon" href="https://cdn.theatlantic.com/assets/static/a/theatlantic/common/img/apple-mask-icon.svg" color="#000000">

        <meta name="application-name" content="">
        <meta name="msapplication-TileColor" content="">
        <meta name="msapplication-TileImage" content="">

        <meta name="apple-itunes-app" content="app-id=397599894">

        
        <link rel="alternate" type="application/rss+xml" title="The Atlantic" href="/feed/all/" />
        <link rel="alternate" type="application/rss+xml" title="Best of The Atlantic" href="/feed/best-of/" />
        
        

        
        
        <meta name="keywords" content="The Atlantic, The Atlantic Magazine, TheAtlantic.com, Atlantic, news, opinion, breaking news, analysis, commentary, business, politics, culture, international, science, technology, national and life">
        <meta name="description" content="Power outages in Venezuela, Mardi Gras festivities in New Orleans, presidential campaigning in New York, International Women’s Day observed around the world, and much more">
        <link rel="image_src" href="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845-1/facebook.jpg?1552069712">

        
        <meta name="referrer" content="unsafe-url" />

        <meta property="og:site_name" content="The Atlantic">
        <meta property="og:locale" content="en_US">
        <meta property="fb:admins" content="577048155,17301937">
        <meta property="fb:app_id" content="100770816677686">

        
        <meta property="fb:pages" content="29259828486,1468531833474495,1061579677251147,457711054591520,370457103090695,1631141167169115,148681772342453,1510507419185410,128344747344340,128377530562508,236061986423933" />

        <link href="https://plus.google.com/109258622984321091629" rel="publisher" />
        <meta name="p:domain_verify" content="68e1a0361a557708fefc992f3309ed70"/>

        <meta property="og:title" content="Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set">
        <meta property="og:type" content="article">
        <meta property="og:url" content="https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/" />

        
            <link rel="canonical" href="https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/" />
        

        <meta property="og:image" content="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845-1/facebook.jpg?1552069712">
        <meta property="og:description" content="Power outages in Venezuela, Mardi Gras festivities in New Orleans, presidential campaigning in New York, International Women&#39;s Day observed around the world, and much more">

        
        
            
        
        <meta name="author" content="Alan Taylor" />
        

        <meta name="twitter:site" content="@theatlantic" />
        <meta name="twitter:domain" content="theatlantic.com" />
        
<meta name="twitter:card" content="gallery" />

<meta name="twitter:image0" content="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845/main_1500.jpg">

<meta name="twitter:image1" content="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w02_1129218742/main_1500.jpg">

<meta name="twitter:image2" content="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w03_1128958183/main_1500.jpg">

<meta name="twitter:image3" content="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w04_1128496602/main_1500.jpg">



        
            
                <meta name="ROBOTS" content="INDEX, FOLLOW">
            
        

    

    
        <script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "Organization",
  "name": "The Atlantic",
  "url": "https://www.theatlantic.com",
  "logo": "https://cdn.theatlantic.com/assets/media/files/atlantic-logo--224x224.png",
  "sameAs" : [
    "https://www.facebook.com/TheAtlantic",
    "https://twitter.com/theatlantic",
    "https://plus.google.com/109258622984321091629/posts"]
}
</script>
<script type="application/ld+json">
{
  "@context" : "http://schema.org",
  "@type" : "WebSite",
  "name" : "The Atlantic",
  "url" : "https://www.theatlantic.com",
  "potentialAction" : {
    "@type" : "SearchAction",
    "target" : "https://www.theatlantic.com/search/?q={search_term}",
    "query-input" : "required name=search_term"
  }
}
</script>

        
            





<script type="application/ld+json">
{
  "@context": "https://schema.org/",
  "@type": "NewsArticle",
  "headline": "Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set",
  "url": "https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/",
  
  "image": "https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845-1/original.jpg",
  
  
  "thumbnailUrl": "https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845-1/thumb_wide.jpg?1552069712",
  
  "datePublished": "2019-03-08T14:26:13-0500",
  "dateModified": "2019-03-08T16:16:37-0500",
  "articleSection": "Global",
  "mainEntityOfPage": {
    "@type": "WebPage",
    "@id": "https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/"
  },
  "publisher": {
    "@type": "Organization",
    "name": "The Atlantic",
    "url": "https://www.theatlantic.com",
    "logo": {
      "@type": "ImageObject",
      "url": "https://cdn.theatlantic.com/assets/media/files/atlantic-logo--224x224.png"
    }
  },
  "author": [
      {
        "@type": "Person",
        "name": "Alan Taylor",
        "sameAs": "https://www.theatlantic.com/author/alan-taylor/"
      }
  ],
  "creator": [
      {
        "@type": "Person",
        "name": "Alan Taylor",
        "sameAs": "https://www.theatlantic.com/author/alan-taylor/"
      }
  ]
}
</script>




        
    

    
<link rel="amphtml" href="https://www.theatlantic.com/amp/photo/584431/" />
<meta property="ia:markup_url" content="https://www.theatlantic.com/facebook-instant/photo/584431/">


    
    <style>@font-face{font-display:swap;font-family:Druk;font-style:normal;font-weight:800;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/druk/Druk-Heavy-Web.73c680315608.woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/druk/Druk-Heavy-Web.de0a29cd8676.woff")format("woff")}@font-face{font-display:swap;font-family:"Graphik";font-style:normal;font-weight:400;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/graphik/Graphik-Regular-Web.f00835c4ea93.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/graphik/Graphik-Regular-Web.475a966967ba.woff")format("woff")}
@font-face{font-display:swap;font-family:"Graphik";font-style:normal;font-weight:600;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/graphik/Graphik-Semibold-Web.4ddb0b7ae595.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/graphik/Graphik-Semibold-Web.86883f41da8f.woff")format("woff")}@font-face{font-display:swap;font-family:"Lyon Text";font-stretch:normal;font-style:normal;font-weight:400;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Regular.952052f78892.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Regular.840c78933255.woff")format("woff")}
@font-face{font-display:swap;font-family:"Lyon Text";font-stretch:normal;font-style:italic;font-weight:400;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Regular-Italic.1f3d1b40255e.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Regular-Italic.ce9636f76967.woff")format("woff")}
@font-face{font-display:swap;font-family:"Lyon Text";font-stretch:normal;font-style:normal;font-weight:700;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Bold.8f7a8108a163.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Bold.49ef3785215a.woff")format("woff")}
@font-face{font-display:swap;font-family:"Lyon Text";font-stretch:normal;font-style:italic;font-weight:700;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Bold-Italic.22e4aa94ac93.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/lyon/Lyon-Text-Bold-Italic.65ae5cd7ad2d.woff")format("woff")}@font-face{font-display:swap;font-family:"Noe Text";font-stretch:normal;font-style:normal;font-weight:900;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/noe/Noe-Text-Black.b05d72fe331d.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/noe/Noe-Text-Black.2dcd42be9f8e.woff")format("woff")}
@font-face{font-display:swap;font-family:"Noe Text";font-stretch:normal;font-style:italic;font-weight:900;src:url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/noe/Noe-Text-Black-Italic.3e111503a215.woff2")format("woff2"),url("https://cdn.theatlantic.com/assets/static/a/frontend/dist/theatlantic/fonts/noe/Noe-Text-Black-Italic.0086b1b36dfc.woff")format("woff")}</style>

    
    <link href="https://cdn.theatlantic.com/assets/static/a/theatlantic/css/site.min.dd789925099e.css" rel="stylesheet" type="text/css" />

    
<link href="https://cdn.theatlantic.com/assets/static/a/theatlantic/css/photo.7a5db34c8937.css" rel="stylesheet" type="text/css" />


    
        <script id="page_meta" type="application/insights+json">
          {"adtargeting": {"grapeshot_segments": ["an_bankofamerica_negative", "an_facebook_safe", "an_fidelity_safe", "an_jpmc_safe", "an_ms_safe", "gs_home_pets", "gs_sport", "gs_home", "gs_politics_misc", "gs_event_blackfriday", "gs_politics"], "watson_categories": ["/pets/dogs", "/sports/snowboarding"]}, "grapeshot_last_updated": "2019-03-09T21:27:27.728059+00:00", "watson": {"top_category": {"full": "/pets/dogs", "one": "pets", "two": "dogs", "three": null, "four": null}, "categories": ["/pets/dogs", "/sports/snowboarding"], "concepts": ["Major League Baseball", "Mardi Gras", "Carnival", "United States", "National Basketball Association", "U.S. state", "National Football League", "Utah"], "entities": [], "keywords": ["Power outages", "Mardi Gras festivities", "New York", "Crufts dog show", "Carnival celebrations", "presidential campaigning", "Venezuela", "Colorado", "Brazil", "England", "baseball spring training", "Arizona", "world"]}, "watson_last_updated": "2019-03-08T21:21:58.149195+00:00", "meta": {"top_category_full": "/pets/dogs", "top_category_tier_1": "pets", "top_category_tier_2": "dogs", "top_category_tier_3": null, "top_category_tier_4": null, "primary_channel": "international", "editorial_categories": "photos of the week", "view": "amp", "report": null, "authors": "alan taylor", "primary_author": "alan taylor", "secondary_author": null, "article_id": 584431, "content_headline": "photos of the week: frozen road, broadway gorilla, penguin swing set", "content_headline_capitalized": "Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set", "content_headline_character_count": 68, "content_headline_word_count": 11, "first_published_date": "2019-03-08", "first_published_time": "14:26", "first_published_datetime": "2019-03-08T14:26:13-05:00", "last_updated_datetime": "2019-03-08T16:16:37-05:00", "is_404": false, "social_title": "photos of the week: frozen road, broadway gorilla, penguin swing set", "social_description": "power outages in venezuela,\u00a0mardi gras festivities in new orleans, presidential campaigning in new york, international women's day observed around the world, and much more", "content_article_word_count": 46, "content_platform": "web", "grapeshot_segments": "an_bankofamerica_negative,an_facebook_safe,an_fidelity_safe,an_jpmc_safe,an_ms_safe,gs_home_pets,gs_sport,gs_home,gs_politics_misc,gs_event_blackfriday,gs_politics", "watson_categories": "/pets/dogs;/sports/snowboarding", "editorial_project": null, "channels": "international", "content_subtype": "photo"}, "article_word_count": 46, "article_last_updated": "2019-03-08T16:16:37-05:00"}
        </script>
    

    <script>
        var Atlantic = {};

        
            
                Atlantic.page_meta = (function(){
                    var page_meta = document.getElementById('page_meta');
                    var data = {};
                    try {
                        data = JSON.parse(page_meta.textContent);
                    } catch (err) {}
                    return data;
                })();
            
        
        Atlantic.page_info = {"domain": "www.theatlantic.com", "image": "https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845-1/facebook.jpg?1552069712", "channels": ["international"], "article_type": "photo", "is_features": false, "trump_channel": "photo", "twitter_image": "", "canonical_url": "https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/", "date": "2019-03-08T14:26:13-05:00", "is_amp": false, "is_freelance": false, "original_url": "https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/", "seo_title": "Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set", "site_url": "https://www.theatlantic.com", "primary_channel": "international", "version": "1.0.0", "is_accounts": false, "share_dek": "Power outages in Venezuela,\u00a0Mardi Gras festivities in New Orleans, presidential campaigning in New York, International Women's Day observed around the world, and much more", "is_instant": false, "share_title": "Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set", "description": "Power outages in Venezuela, Mardi Gras festivities in New Orleans, presidential campaigning in New York, International Women\u2019s Day observed around the world, and much more", "tags": [], "is_404": false, "kicker": "", "authors": ["Alan Taylor"], "hours_since_published": 0, "path": "/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/", "article_id": 584431, "share_text": "Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set\u2014via @TheAtlPhoto", "url": "https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/", "is_sponsored": false, "title": "Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set", "request": {"experiment_group": null, "experiment_id": null, "is_gdpr": false}, "og_type": "article", "report": "", "is_magazine": false, "view": "article", "days_since_published": 0};
        Atlantic.STATIC_URL = "https://cdn.theatlantic.com/assets/static/a/";
        Atlantic.MEDIA_URL = "https://cdn.theatlantic.com/assets/media/";
        Atlantic.PROFILES_STATICFILES = {"CSS": ["https://cdn.theatlantic.com/assets/static/a/theatlantic/css/profiles.min.9a59d93196f8.css"], "JS": "https://cdn.theatlantic.com/assets/static/a/janrain/profiles.d466b109a086.js"};
        Atlantic.PROFILES_URL = "https://accounts.theatlantic.com/"
        Atlantic.SESSION_URLS = {
            "session_start": "https://accounts.theatlantic.com/api/1.0/session_start/",
            "session_end": "https://accounts.theatlantic.com/api/1.0/session_end/"
        };
        Atlantic.GTM = {
            quantcastLabels: ["The Atlantic.Editorial.Title.Photos of the Week: Frozen Road Broadway Gorilla Penguin Swing Set", "The Atlantic.Editorial.Author.Alan Taylor", "The Atlantic.Editorial.Channel.international"],
            simplereach: {"is_404": false, "conf": {"channels": ["international"], "date": "2019-03-08T14:26:13-05:00", "title": "Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set", "url": "https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/", "authors": ["alan-taylor"], "tags": [], "page_url": "https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/"}}
        };
    </script>
    

    

<link rel="preload" as="script" href="https://www.googletagservices.com/tag/js/gpt.js" onload="performance.mark('atl:gpt_preloaded');">

<link rel="preload" as="script" href="https://z.moatads.com/thewire746082389130/moatad.js" onload="performance.mark('atl:moat_preloaded');">

    
        <link rel="preload" as="script" href="https://www.theatlantic.com/packages/adsjs/prebid.js" onload="performance.mark('atl:prebid_preloaded');">
    

    
        <link rel="preload" as="script" href="https://c.amazon-adsystem.com/aax2/apstag.js" onload="performance.mark('atl:amazon_preloaded');">
    

    
        <link rel="preload" as="script" href="https://static.criteo.net/js/ld/publishertag.js" onload="performance.mark('atl:criteo_preloaded');">
    




    <script type="text/javascript" src="https://cdn.theatlantic.com/assets/static/a/theatlantic/js/head.min.bf906d90a783.js" charset="utf-8"></script>
    





    

<script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
})(window,document,'script','dataLayer','GTM-56LJR35');</script>




    

<script type="text/javascript">
    Modernizr.addTest('fullscreen',function(){
        var ancelFullScreen = 'ancelFullScreen'; //make string minifiable

        var limit = Modernizr._domPrefixes.length;
        for(var i = 0; i < limit; ++i) {
        if( document[Modernizr._domPrefixes[i].toLowerCase() + 'C' + ancelFullScreen])
        return true;
        }
        return !!document['c' + ancelFullScreen] || false;
    });
</script>


    <script src="https://cdn.theatlantic.com/assets/static/a/frontend/jspm_packages/system.822e0a71e40b.js"></script>
</head>

<body
    id="photo-channel"
    class=""
    data-page-type="article"
    
>


<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-56LJR35" height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>



<div class="u-hidden"><?xml version="1.0" encoding="UTF-8"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg xmlns="http://www.w3.org/2000/svg"><symbol id="icon-arrow-right" viewBox="0 0 11 18"><path d="M1 1l8 7.928L1.267 17" fill-rule="evenodd"/></symbol><symbol id="icon-logo-a" viewBox="0 0 15 17"><path d="M11.435 3.13v5.913H7.609l3.826-5.913zM12.826 0L2.391 15.652H1V16h3.826v-.348H3.087l4.174-6.26h4.174v6.26h-1.74V16h4.522v-.348h-1.39V0z" fill-rule="evenodd"/></symbol><symbol id="icon-logo-atl" viewBox="0 0 82 25"><g><path d="M25.509.349h-.669L9.297 23.188H7.976l-.21.845h4.406l.212-.845H10.36l5.62-8.258h5.152l-.934 8.258h-1.803l-.212.845h6.185l.208-.845h-1.642L25.509.349zm-3.35 5.503l-.928 8.205h-4.656l5.584-8.205zM70.052 8.578c.812 0 1.423-.662 1.423-1.54 0-.811-.562-1.422-1.308-1.422-.825 0-1.424.659-1.424 1.567 0 .795.563 1.395 1.309 1.395zM29.809 20.791c-.421.962-1.131 2.084-1.747 2.084-.174 0-.352-.052-.352-.438 0-.2.053-.521.137-.823l2.641-10.29h2.205l.212-.842h-2.202l1.145-4.409H29.27l-1.145 4.409h-1.689l-.212.842h1.687l-2.408 9.445-.048.18c-.135.501-.275 1.02-.275 1.688 0 .764.485 1.653 1.853 1.653 1.438 0 2.605-1.028 3.467-3.058l.165-.359-.759-.305-.097.223zM35.675 20.797c-.279.678-1.073 2.078-1.744 2.078-.207 0-.323-.021-.323-.438 0-.287.031-.525.106-.823L39.118.807h-4.497l-.21.843h1.858l-4.9 19.124-.047.188c-.136.534-.276 1.086-.276 1.68 0 .799.487 1.651 1.853 1.651 2.263 0 3.335-2.664 3.527-3.198l.081-.227-.739-.299-.093.228zM77.5 13.191c0 .671.513 1.193 1.167 1.193.777 0 1.279-.625 1.279-1.593 0-1.199-.792-2.599-3.025-2.599-2.252 0-3.8 1.357-4.731 4.147-.444 1.311-1.272 3.747-1.272 6.406 0 1.712.908 3.542 3.456 3.542 1.238 0 2.98-.668 4.25-3.854l.14-.33-.829-.381-.1.238c-.32.762-1.293 3.081-2.889 3.081-.974 0-1.468-.744-1.468-2.211 0-2.229.673-4.556 1.332-6.534.797-2.386 1.245-3.233 2.168-3.233.56 0 .951.226.951.549 0 .284-.097.497-.2.723-.107.242-.229.51-.229.856zM70.573 10.48H64.79l1.117-4.409h-2.576l-1.117 4.409h-1.889l-.209.842h1.886l-2.41 9.447-.045.172c-.136.513-.277 1.041-.277 1.638 0 .789.5 1.71 1.91 1.71 2.166 0 3.246-2.664 3.441-3.197l.083-.228-.736-.296-.098.222c-.57 1.305-1.223 2.084-1.746 2.084-.189 0-.323-.031-.323-.438 0-.229.041-.475.136-.822l2.615-10.289 3.258.001-2.437 9.442c-.137.492-.324 1.168-.324 1.756 0 .814.507 1.768 1.938 1.768 2.157 0 3.245-2.663 3.441-3.197l.083-.228-.741-.274-.1.215c-.84 1.803-1.432 2.069-1.71 2.069-.281 0-.352-.055-.352-.466 0-.254.049-.483.135-.797l2.83-11.134zM58.158 20.88c-.27.546-.988 1.995-1.674 1.995-.297 0-.378-.088-.378-.408 0-.238.045-.523.134-.854l1.832-7.101.047-.18c.156-.593.304-1.151.304-1.886 0-1.094-.759-2.197-2.454-2.197-1.026 0-1.937.428-2.718 1.275l.261-1.045H49.09l-.279.842h1.884L47.431 24.03h2.577l2.281-8.802c.495-1.871 1.633-3.765 2.851-3.765.692 0 .837.44.837.809 0 .363-.102.762-.22 1.224l-1.872 7.306c-.117.477-.314 1.271-.314 1.83 0 .799.502 1.653 1.911 1.653.958 0 2.33-.538 3.428-3.103l.143-.315-.743-.3-.152.313zM47.765 20.866l-.743-.298-.097.223c-.789 1.812-1.432 2.083-1.745 2.083-.164 0-.217-.058-.235-.074-.069-.073-.101-.229-.088-.448 0-.158.045-.404.135-.737l2.857-11.133h-2.58l-.202.828c-.418-.74-1.128-1.114-2.12-1.114-2.718 0-3.95 2.378-4.641 4.167-.39 1.022-1.415 4.151-1.477 6.32-.047 1.209.25 2.183.857 2.812.5.521 1.201.796 2.026.796 1.22 0 2.087-.83 2.617-1.567.03.389.186.743.449 1.018.348.361.851.552 1.456.552 1.436 0 2.65-1.116 3.439-3.19l.092-.238zm-7.564 2.238c-.312 0-.538-.076-.69-.229-.208-.213-.305-.606-.291-1.171.072-2.161 1.23-6.062 1.904-8.028.385-1.136 1.073-2.491 2.167-2.491h.001c.299 0 .546.099.735.293.261.269.402.715.388 1.222-.053 1.295-.642 3.936-1.535 6.888-.692 2.2-1.693 3.516-2.679 3.516zM6.992 3.046l-.052.294h.721L8.322.75H2.71l-.658 2.59h.719l.062-.158c.582-1.478.76-1.665 1.593-1.665h.132L2.519 9.494h-.832l-.215.767h3.33l.225-.767H4.1l2.041-7.977h.186c.383 0 .622.02.707.121.153.184.058.849-.042 1.408zM9.119 6.277c.257-.957.775-1.679 1.204-1.679.163 0 .245.069.245.206 0 .204-.076.497-.144.736L9.67 8.45c-.039.116-.166.639-.166.954 0 .599.373.972.973.972.729 0 1.31-.506 1.745-1.547.089.996.686 1.546 1.713 1.546 1.378 0 2.041-1.46 2.211-1.908l.081-.211-.666-.336-.105.244c-.202.468-.642 1.253-1.329 1.253-.126 0-.511 0-.511-.699 0-.351.058-.703.113-.973 1.376-.358 3.193-1.427 3.193-2.829 0-.799-.519-1.276-1.389-1.276-1.643 0-3.315 2.465-3.329 4.896l-.528-.217-.097.225c-.248.576-.508.838-.613.876 0-.001-.01-.023-.01-.093 0-.156.069-.507.106-.656l.765-2.872c.076-.309.154-.627.154-.995 0-.578-.395-1.162-1.276-1.162a1.72 1.72 0 0 0-1.002.337l.829-3.228H8.424l-.217.768h.655l-2.243 8.742h1.466l1.034-3.984zm4.83.567c.392-1.402 1.132-2.437 1.585-2.437.184 0 .203.019.203.207.001.742-.715 1.759-1.788 2.23z"/></g></symbol><symbol id="icon-nav-latest" viewBox="0 0 24 24"><path d="M12 24c6.627 0 12-5.373 12-12S18.627 0 12 0 0 5.373 0 12s5.373 12 12 12zm0-2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm1-18.003a1 1 0 1 0-2 .006l.024 9c.002.893 1.085 1.335 1.712.7l3.976-4.034a1 1 0 0 0-1.424-1.404l-2.27 2.303L13 3.998z" fill-rule="evenodd"/></symbol><symbol id="icon-nav-magazine" viewBox="0 0 24 23"><path d="M22 19.108l-9 1.03V3.163l9-1.03v16.974zM2 2.134l9 1.03v16.973l-9-1.029V2.134zM23.001.006L12 1.266.999.005C.447-.057 0 .337 0 .887v19.01c0 .544.447 1.046.999 1.11l10.586 1.21a.928.928 0 0 0 .414.048h.002a.928.928 0 0 0 .414-.048l10.586-1.21c.552-.064.999-.566.999-1.11V.887c0-.55-.447-.944-.999-.88z" fill-rule="evenodd"/></symbol><symbol id="icon-nav-more" viewBox="0 0 22 22"><path d="M10 10V1a1 1 0 0 1 2 0v9h9a1 1 0 0 1 0 2h-9v9a1 1 0 0 1-2 0v-9H1a1 1 0 0 1 0-2h9z" fill-rule="evenodd"/></symbol><symbol id="icon-nav-popular" viewBox="0 0 24 22"><path d="M20.297 3.566l-7.198 12.602a1 1 0 0 1-1.695.068l-2.9-4.251-6.693 9.27a1 1 0 1 1-1.622-1.172L7.716 9.66a1 1 0 0 1 1.637.022l2.802 4.107L18.59 2.52l-4.365 1.007a1 1 0 0 1-.45-1.948l6.73-1.554a1 1 0 0 1 1.2.75l1.596 6.914a1 1 0 1 1-1.948.45l-1.056-4.574z" fill-rule="evenodd"/></symbol><symbol id="icon-nav-sections" viewBox="0 0 20 20"><path d="M1 20h18a1 1 0 0 0 0-2H1a1 1 0 0 0 0 2zm0-6h18a1 1 0 0 0 0-2H1a1 1 0 0 0 0 2zm0-6h18a1 1 0 0 0 0-2H1a1 1 0 1 0 0 2zm0-6h18a1 1 0 0 0 0-2H1a1 1 0 1 0 0 2z" fill-rule="evenodd"/></symbol><symbol id="icon-nav-writers" viewBox="0 0 24 23"><path d="M2.013 21.042l3.288.041h16.663c-.279-2.07-1.549-6.83-7.332-7.951a.956.956 0 0 1-.272-1.785c1.611-.865 2.613-2.564 2.613-4.435 0-2.755-2.174-4.995-4.845-4.995s-4.844 2.24-4.844 4.995c0 1.836.974 3.52 2.54 4.395a.957.957 0 0 1-.3 1.781c-5.827 1.03-7.192 5.857-7.51 7.954M23 23H5.289l-4.342-.055A.963.963 0 0 1 0 21.97c.002-.085.245-7.823 7.252-10.268a6.995 6.995 0 0 1-1.885-4.79C5.367 3.102 8.401 0 12.128 0c3.728 0 6.762 3.1 6.762 6.912 0 1.859-.721 3.596-1.961 4.87 6.942 2.543 7.03 10.176 7.03 10.26 0 .53-.43.958-.96.958" fill-rule="evenodd"/></symbol><symbol id="icon-social-email" viewBox="0 0 21 16"><path d="M21 2.254l-10.5 7.72L0 2.196A2 2 0 0 1 2 .222h17a2 2 0 0 1 2 2v.032zm0 2.369v9.155a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-9.21c6.081 4.581 9.297 7.004 9.65 7.268.477.359.591.428.855.426.258-.001.372-.07.834-.415L21 4.623z" fill-rule="evenodd"/></symbol><symbol id="icon-social-facebook" viewBox="0 0 16 16"><path d="M15.117 0H.883A.883.883 0 0 0 0 .883v14.234c0 .488.395.883.883.883h7.663V9.804H6.461V7.389h2.085V5.61c0-2.067 1.262-3.192 3.106-3.192.883 0 1.642.065 1.863.095v2.16h-1.279c-1.002 0-1.196.476-1.196 1.176v1.541h2.39l-.31 2.415h-2.08V16h4.077a.883.883 0 0 0 .883-.883V.883A.883.883 0 0 0 15.117 0" fill-rule="evenodd"/></symbol><symbol id="icon-social-flipboard" viewBox="0 0 500 500"><path d="M0 0v500h500V0H0zm400 200H300v100H200v100H100V100h300v100z"/></symbol><symbol id="icon-social-google" viewBox="0 0 16 16"><path d="M8.163 6.857V9.6h4.63c-.187 1.177-1.4 3.451-4.63 3.451-2.787 0-5.061-2.262-5.061-5.051s2.274-5.051 5.061-5.051c1.586 0 2.647.662 3.254 1.234l2.216-2.092C12.21.79 10.367 0 8.163 0 3.65 0 0 3.577 0 8s3.65 8 8.163 8C12.875 16 16 12.754 16 8.183c0-.526-.058-.926-.128-1.326H8.163z" fill-rule="evenodd"/></symbol><symbol id="icon-social-instagram" viewBox="0 0 16 16"><path d="M8 0C5.827 0 5.555.01 4.702.048 3.85.087 3.269.222 2.76.42a3.921 3.921 0 0 0-1.417.923c-.445.444-.719.89-.923 1.417-.198.509-.333 1.09-.372 1.942C.01 5.555 0 5.827 0 8s.01 2.445.048 3.298c.039.852.174 1.433.372 1.942.204.526.478.973.923 1.417.444.445.89.719 1.417.923.509.198 1.09.333 1.942.372C5.555 15.99 5.827 16 8 16s2.445-.01 3.298-.048c.852-.039 1.433-.174 1.942-.372a3.921 3.921 0 0 0 1.417-.923c.445-.444.719-.89.923-1.417.198-.509.333-1.09.372-1.942C15.99 10.445 16 10.173 16 8s-.01-2.445-.048-3.298c-.039-.852-.174-1.433-.372-1.942a3.921 3.921 0 0 0-.923-1.417A3.921 3.921 0 0 0 13.24.42c-.509-.198-1.09-.333-1.942-.372C10.445.01 10.173 0 8 0zm0 1.442c2.136 0 2.39.008 3.233.046.78.036 1.203.166 1.485.276.374.145.64.318.92.598.28.28.453.546.598.92.11.282.24.705.276 1.485.038.844.046 1.097.046 3.233s-.008 2.39-.046 3.233c-.036.78-.166 1.203-.276 1.485-.145.374-.318.64-.598.92-.28.28-.546.453-.92.598-.282.11-.705.24-1.485.276-.844.038-1.097.047-3.233.047s-2.39-.009-3.233-.047c-.78-.036-1.203-.166-1.485-.276a2.478 2.478 0 0 1-.92-.598 2.478 2.478 0 0 1-.598-.92c-.11-.282-.24-.705-.276-1.485-.038-.844-.047-1.097-.047-3.233s.009-2.39.047-3.233c.036-.78.166-1.203.276-1.485.145-.374.318-.64.598-.92.28-.28.546-.453.92-.598.282-.11.705-.24 1.485-.276.844-.038 1.097-.046 3.233-.046zM7.9 3.8a4.1 4.1 0 1 0 0 8.2 4.1 4.1 0 0 0 0-8.2zm0 6.761a2.661 2.661 0 1 1 0-5.322 2.661 2.661 0 0 1 0 5.322zM13.4 3.8a1 1 0 1 1-2 0 1 1 0 0 1 2 0z" fill-rule="evenodd"/></symbol><symbol id="icon-social-linkedin" viewBox="0 0 16 16"><path d="M3.502 15.925H.343V5.085h3.09v10.84h.07zM1.854 3.738C.687 3.738 0 2.916 0 1.87 0 .822.755 0 1.923 0 3.09 0 3.777.822 3.777 1.87c0 .971-.756 1.868-1.923 1.868zM16 15.925h-3.57v-5.607c0-1.496-.55-2.468-1.786-2.468-.962 0-1.442.673-1.717 1.346-.069.225-.069.598-.069.897V16H5.356s.069-9.944 0-10.916h3.502v1.72c.206-.748 1.305-1.795 3.09-1.795 2.198 0 3.983 1.57 3.983 4.935v5.981H16z" fill-rule="evenodd"/></symbol><symbol id="icon-social-pinterest" viewBox="0 0 16 16"><path d="M16 8c0-4.417-3.583-8-8-8a8.002 8.002 0 0 0-3.208 15.333c-.021-.562-.01-1.24.135-1.843 0 0 .156-.646 1.031-4.355-.26-.51-.26-1.27-.26-1.27 0-1.188.687-2.073 1.542-2.073.729 0 1.083.552 1.083 1.208 0 .73-.469 1.823-.708 2.833-.198.855.427 1.542 1.26 1.542 1.52 0 2.542-1.948 2.542-4.25 0-1.76-1.188-3.073-3.334-3.073-2.427 0-3.937 1.813-3.937 3.833 0 .698.208 1.188.531 1.573.146.177.167.24.115.448-.042.146-.125.5-.167.636-.052.208-.219.28-.396.208-1.114-.458-1.635-1.687-1.635-3.062 0-2.271 1.916-5 5.708-5 3.063 0 5.073 2.218 5.073 4.583 0 3.146-1.75 5.49-4.323 5.49-.865 0-1.677-.47-1.948-1a345.426 345.426 0 0 1-.562 2.197c-.167.615-.51 1.23-.813 1.709C6.46 15.885 7.21 16 8 16c4.417 0 8-3.583 8-8z" fill-rule="evenodd"/></symbol><symbol id="icon-social-rss" viewBox="0 0 16 16"><path d="M2.74 16a2.74 2.74 0 1 0 0-5.48 2.74 2.74 0 0 0 0 5.48zm11.89 0a1.37 1.37 0 0 1-1.37-1.37c0-7.335-4.556-11.89-11.89-11.89a1.37 1.37 0 0 1 0-2.74c4.331 0 8 1.384 10.625 4.005C14.62 6.625 16 10.3 16 14.63A1.37 1.37 0 0 1 14.63 16zm-5.185 0a1.37 1.37 0 0 1-1.37-1.37c0-4.449-2.256-6.705-6.705-6.705a1.37 1.37 0 0 1 0-2.74c6.001 0 9.445 3.442 9.445 9.444A1.37 1.37 0 0 1 9.445 16z" fill-rule="evenodd"/></symbol><symbol id="icon-social-twitter" viewBox="0 0 16 14"><path d="M14.115 2.13A3.393 3.393 0 0 0 15.558.25a6.437 6.437 0 0 1-2.085.825A3.226 3.226 0 0 0 11.077 0C9.265 0 7.795 1.523 7.795 3.401c0 .267.029.527.085.776C5.152 4.035 2.733 2.68 1.114.623a3.48 3.48 0 0 0-.445 1.71c0 1.18.58 2.22 1.46 2.83a3.183 3.183 0 0 1-1.486-.425v.043c0 1.648 1.131 3.023 2.633 3.335a3.18 3.18 0 0 1-1.483.059c.418 1.351 1.63 2.334 3.067 2.362A6.442 6.442 0 0 1 0 11.946a9.057 9.057 0 0 0 5.032 1.528c6.038 0 9.34-5.183 9.34-9.678 0-.148-.004-.294-.01-.44A6.811 6.811 0 0 0 16 1.595a6.385 6.385 0 0 1-1.885.536z" fill-rule="evenodd"/></symbol><symbol id="icon-social-youtube" viewBox="0 0 16 12"><path d="M11.428 6a.589.589 0 0 1-.267.506l-4.572 3a.508.508 0 0 1-.303.094.585.585 0 0 1-.277-.075A.613.613 0 0 1 5.714 9V3c0-.216.116-.422.295-.525a.542.542 0 0 1 .58.019l4.572 3c.17.103.268.3.268.506zM16 6c0-1.34 0-2.766-.277-4.069-.196-.919-.893-1.594-1.732-1.697C12.01 0 10 0 8 0 6 0 3.991 0 2.009.234 1.169.338.482 1.013.286 1.931 0 3.234 0 4.66 0 6c0 1.34 0 2.766.277 4.069.196.919.893 1.594 1.732 1.697C3.99 12 6 12 8 12c2 0 4.009 0 5.991-.234.84-.104 1.536-.779 1.723-1.697C16 8.766 16 7.34 16 6z" fill-rule="evenodd"/></symbol><symbol id="icon-ui-ellipses" viewBox="0 0 16 4"><g fill-rule="evenodd"><circle cx="2" cy="2" r="2"/><circle cx="8" cy="2" r="2"/><circle cx="14" cy="2" r="2"/></g></symbol><symbol id="icon-ui-lock" viewBox="0 0 12 14"><path d="M10 5h1.5a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H.5a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 .5 5H2V4a4 4 0 0 1 4-4 4 4 0 0 1 4 4zM8 5V4a2 2 0 0 0-2-2 2 2 0 0 0-2 2v1z"/></symbol><symbol id="icon-ui-search" viewBox="0 0 16 16"><path d="M16 14.783l-4.896-4.83c.866-1.065 1.237-2.28 1.237-3.763C12.34 2.771 9.578 0 6.17 0 2.762 0 0 2.771 0 6.19c0 3.419 2.763 6.19 6.17 6.19 1.486 0 2.613-.33 3.678-1.207L14.742 16 16 14.783zM1.234 6.189c0-2.735 2.21-4.953 4.936-4.953 2.727 0 4.937 2.218 4.937 4.953a4.943 4.943 0 0 1-4.937 4.952c-2.726 0-4.936-2.217-4.936-4.952z" fill-rule="evenodd"/></symbol><symbol id="icon-ui-x" viewBox="0 0 16 16"><path d="M9.525 8l6.159 6.159a1.078 1.078 0 1 1-1.525 1.525L8 9.524l-6.159 6.16a1.076 1.076 0 0 1-1.525 0 1.078 1.078 0 0 1 0-1.525L6.476 8 .315 1.841A1.078 1.078 0 1 1 1.841.316L8 6.476l6.16-6.16a1.078 1.078 0 1 1 1.524 1.525L9.524 8z" fill-rule="evenodd"/></symbol></svg></div>



<div id="site" class="">
    

        
        <nav id="main-navigation" class="c-nav c-nav--main u-border-box u-smooth-font" role="navigation" aria-label="Main menu"><div class="c-nav__container"><ul itemscope itemtype="http://www.schema.org/SiteNavigationElement" class="c-nav__list c-nav__list--main" role="menubar" aria-hidden="false" id="main-navigation-top-level"><li class="c-nav__item c-nav__item--main c-nav__item--logo"> <a href="#main-content" class="c-skip-link">Skip to content</a> <a href="/" class="c-nav__link c-nav__link--main c-logo c-logo--nav"><span class="u-element-invisible c-logo__name">The Atlantic</span> <svg class="o-icon c-logo__icon"  viewBox="0 0 256 84"  ><use xlink:href="#icon-logo-atl"/></svg></a> </li><li class="c-nav__item c-nav__item--main c-nav__item--popular js-nav-item" role="menuitem"><a itemprop="url" href="/most-popular/" class="c-nav__link c-nav__link--main" id="main-navigation-popular"><svg class="o-icon c-nav__icon"  ><use xlink:href="#icon-nav-popular"/></svg> <span itemprop="name" class="c-nav__title c-nav__title--popular">Popular</span></a></li><li class="c-nav__item c-nav__item--main js-nav-item" role="menuitem"><a itemprop="url" href="/latest/" class="c-nav__link c-nav__link--main c-nav__link--latest" id="main-navigation-latest"><svg class="o-icon c-nav__icon"  ><use xlink:href="#icon-nav-latest"/></svg> <span itemprop="name" class="c-nav__title c-nav__title--latest">Latest</span></a></li>  <li itemscope itemtype="http://www.schema.org/SiteNavigationElement" class="c-nav__item c-nav__item--main has-children js-nav-item" role="menuitem" aria-haspopup="true"><a href="javascript:void(0)" class="c-nav__link c-nav__link--main c-nav__link--sections has-children js-expandable-menu" id="main-navigation-sections"><svg class="o-icon c-nav__icon c-nav__icon--sections"  ><use xlink:href="#icon-nav-sections"/></svg> <span class="c-nav__title c-nav__title--sections">Sections</span></a><div class="c-nav__menu c-menu c-menu--expandable" role="menu" aria-hidden="true"><strong class="c-menu__title"><span class="c-menu__name">Sections</span> <svg class="o-icon c-menu__close c-icon--sections"  ><use xlink:href="#icon-ui-x"/></svg></strong><section class="c-menu__section" id="main-navigation-sections-menu-0"><ul class="c-menu__list" aria-hidden="true" role="menu">  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/politics/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-0">Politics</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/entertainment/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-1">Culture</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/technology/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-2">Technology</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/ideas/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-3">Ideas</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/science/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-4">Science</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/books/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-5">Books</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/family/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-6">Family</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/business/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-7">Business</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/international/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-8">Global</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/health/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-9">Health</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/education/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-10">Education</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/letters/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-11">Letters</a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/membership/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-0-item-12">The Masthead</a></li>  </ul></section><section class="c-menu__section c-menu__section--secondary" id="main-navigation-sections-menu-1"><ul itemscope itemtype="http://www.schema.org/SiteNavigationElement" class="c-menu__list c-menu__list--secondary" aria-hidden="true" role="menu">  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/photo/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-1-item-0">Photo</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/podcasts/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-1-item-1">Podcasts</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/free-daily-crossword-puzzle/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-1-item-2">The Atlantic Crossword</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/video/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-1-item-3">Video</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/events/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-1-item-4">Events</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/writers/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-1-item-5">Writers</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/projects/" class="c-menu__link" tabindex="-1" id="main-navigation-sections-menu-1-item-6">Projects</a></li>  </ul></section></div></li><li itemscope itemtype="http://www.schema.org/SiteNavigationElement" class="c-nav__item c-nav__item--main has-children js-nav-item" role="menuitem" aria-haspopup="true"><a href="javascript:void(0)" class="c-nav__link c-nav__link--main c-nav__link--magazine has-children js-expandable-menu" id="main-navigation-magazine"><svg class="o-icon c-nav__icon c-nav__icon--magazine"  ><use xlink:href="#icon-nav-magazine"/></svg> <span class="c-nav__title c-nav__title--magazine">Magazine</span></a><div class="c-nav__menu c-menu c-menu--expandable" role="menu" aria-hidden="true"><strong class="c-menu__title"><span class="c-menu__name">Magazine</span> <svg class="o-icon c-menu__close c-icon--sections"  ><use xlink:href="#icon-ui-x"/></svg></strong><section class="c-menu__section" id="main-navigation-magazine-menu"><ul class="c-menu__list" aria-hidden="true" role="menu">  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/magazine/" class="c-menu__link" tabindex="-1" id="main-navigation-magazine-menu-item-0"> Current issue </a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/magazine/backissues/" class="c-menu__link" tabindex="-1" id="main-navigation-magazine-menu-item-1"> All issues </a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://accounts.theatlantic.com/" class="c-menu__link" tabindex="-1" id="main-navigation-magazine-menu-item-2"> Manage subscription </a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/giveagift2018/" class="c-menu__link" tabindex="-1" id="main-navigation-magazine-menu-item-3"> Give a Gift </a></li>  <li itemprop="name" class="c-menu__item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/locator/navdropdown/" class="c-menu__link" tabindex="-1" id="main-navigation-magazine-menu-item-4"> Subscribe </a></li>  </ul></section></div></li><li class="c-nav__item c-nav__item--main has-children js-nav-item" role="menuitem" aria-haspopup="true"><a href="javascript:void(0)" class="c-nav__link c-nav__link--main has-children c-nav__link--more js-expandable-menu js-load-janrain" id="main-navigation-more"><svg class="o-icon c-nav__icon c-nav__icon--more"  ><use xlink:href="#icon-nav-more"/></svg> <span class="c-nav__title c-nav__title--more">More <span class="u-element-invisible">Categories</span></span></a><div class="c-nav__menu c-menu c-menu--expandable" role="menu" aria-hidden="true"><strong class="c-menu__title"><span class="c-menu__name">More</span> <svg class="o-icon c-menu__close c-icon--sections"  ><use xlink:href="#icon-ui-x"/></svg></strong><section class="c-menu__section" id="main-navigation-more-menu-0"><ul itemscope itemtype="http://www.schema.org/SiteNavigationElement" class="c-menu__list" aria-hidden="true" role="menu"><li class="c-menu__item"><a itemprop="url" href="https://accounts.theatlantic.com/register/" class="c-menu__link captureSignInLink signed-out-utils" tabindex="-1" id="main-navigation-more-menu-0-item-0"><span itemprop="name">Create account</span></a> <a itemprop="url" href="https://accounts.theatlantic.com/" class="c-menu__link captureProfileLink signed-in-utils" tabindex="-1" id="main-navigation-more-menu-0-item-1"><span itemprop="name">Your account</span></a></li><li class="c-menu__item"><a itemprop="url" href="https://accounts.theatlantic.com/" class="c-menu__link captureSignInLink signed-out-utils" tabindex="-1" id="main-navigation-more-menu-0-item-2"><span itemprop="name">Sign in</span></a> <a itemprop="url" href="javascript:void(0)" class="c-menu__link captureSignOutLink signed-in-utils" tabindex="-1" id="main-navigation-more-menu-0-item-3"><span itemprop="name">Sign out</span></a></li></ul></section><section class="c-menu__section c-menu__section--secondary" id="main-navigation-more-menu-1"><ul itemscope itemtype="http://www.schema.org/SiteNavigationElement" class="c-menu__list c-menu__list--secondary" aria-hidden="true" role="menu">  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/newsletters/sign-up/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-0">Newsletters</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/free-daily-crossword-puzzle/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-1">The Atlantic Crossword</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/app/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-2">iOS App</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/timeline/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-3">Life Timeline</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/events/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-4">Events</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/ebooks/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-5">Books</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/store/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-6">Shop</a></li>  <li itemprop="name" class="c-menu__item c-menu__item--secondary" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/faq/" class="c-menu__link" tabindex="-1" id="main-navigation-more-menu-1-item-7">Help Center</a></li>  </ul></section><section class="c-menu__section c-menu__section--tertiary" aria-hidden="true" role="menu" id="main-navigation-more-menu-2"> <span class="c-menu__section__icon"><a href="https://www.facebook.com/TheAtlantic" target="_blank" rel="noopener" tabindex="-1"  id="main-navigation-more-menu-2-item-0"  >  <svg class="o-icon"  ><use xlink:href="#icon-social-facebook"/></svg>  </a></span>  <span class="c-menu__section__icon"><a href="https://www.instagram.com/theatlantic/" target="_blank" rel="noopener" tabindex="-1"  id="main-navigation-more-menu-2-item-1"  >  <svg class="o-icon"  ><use xlink:href="#icon-social-instagram"/></svg>  </a></span>  <span class="c-menu__section__icon"><a href="https://twitter.com/TheAtlantic" target="_blank" rel="noopener" tabindex="-1"  id="main-navigation-more-menu-2-item-2"  >  <svg class="o-icon"  ><use xlink:href="#icon-social-twitter"/></svg>  </a></span>  <span class="c-menu__section__icon"><a href="http://www.theatlantic.com/follow-the-atlantic/#follow-rssfeeds" target="_blank" rel="noopener" tabindex="-1"  id="main-navigation-more-menu-2-item-3"  >  <svg class="o-icon"  ><use xlink:href="#icon-social-rss"/></svg>  </a></span>  <span class="c-menu__section__icon"><a href="/follow-the-atlantic/" target="_blank" rel="noopener" tabindex="-1"  id="main-navigation-more-menu-2-item-4"  >  View all  </a></span> </section></div></li>  <li class="c-nav__item c-nav__item--main c-nav__item--upper-shelf c-nav__item--subscribe js-nav-item" role="menuitem"><a itemprop="url" href="https://www.theatlantic.com/locator/subscribe-magazine/" class="c-nav__link c-nav__link--main c-nav__link--subscribe"><span itemprop="name" class="c-nav__title c-nav__title--subscribe">Subscribe</span></a></li><li class="c-nav__item c-nav__item--main c-nav__item--upper-shelf c-nav__item--search has-children"> <div class="c-search__container--mobile"><a href="javascript:void(0)" class="c-nav__link has-children js-expandable-menu"><span class="u-element-invisible">Search</span> <svg class="c-search__icon"><use xlink:href="#icon-ui-search"></use></svg></a><div class="c-nav__menu c-menu c-menu--expandable"><strong class="c-menu__title"><span class="c-menu__name">Search</span> <svg class="c-menu__close o-icon"><use xlink:href="#icon-ui-x"></use></svg></strong><section class="c-menu__section" id="main-navigation-search-menu-mobile"><form action="/search/" class="c-search__form c-search__form--mobile"><input aria-label="Search The Atlantic" type="search" name="q" class="c-search c-search--mobile js-search--mobile" placeholder="The Atlantic..."> <input type="submit" class="c-search__submit c-search__submit--mobile"> <svg class="c-search__icon c-search__icon--inline c-search__icon--inline--mobile o-icon"><use xlink:href="#icon-ui-search"></use></svg></form> <strong class="c-menu__title--secondary"> Quick Links </strong>  <ul class="c-menu__list">  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/james-fallows/" class="c-menu__link" tabindex="-1"> James Fallows </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/megan-garber/" class="c-menu__link" tabindex="-1"> Megan Garber </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/david-frum/" class="c-menu__link" tabindex="-1"> David Frum </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/adam-serwer/" class="c-menu__link" tabindex="-1"> Adam Serwer </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="https://accounts.theatlantic.com/" class="c-menu__link" tabindex="-1"> Manage subscription </a></li>  </ul></section></div></div>  <div class="c-search__container__container js-search-container"><div class="c-search__container"><a href="javascript:void(0)" class="js-search-button c-search__open c-nav__search-icon-container"><svg class="c-search__icon"><use xlink:href="#icon-ui-search"></use></svg></a><form class="c-search__form" action="/search/"><input type="submit" class="c-search__submit" tabindex="-1" value="Submit"><label class="u-element-invisible" for="search">Search The Atlantic</label><input type="search" name="q" id="search" class="c-search js-search" placeholder="Search The Atlantic..." autocomplete="off" disabled="disabled"></form><a href="javascript:void(0)" class="js-close c-search__close" tabindex="-1"><svg class="c-search__icon" stroke-width="2"><use xlink:href="#icon-ui-x"></use></svg></a><div class="c-search__menu c-menu is-open"><section class="c-menu__section" id="main-navigation-search-menu-desktop"> <strong class="c-menu__title--secondary"> Quick Links </strong>  <ul class="c-menu__list">  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/james-fallows/" class="c-menu__link" tabindex="-1"> James Fallows </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/megan-garber/" class="c-menu__link" tabindex="-1"> Megan Garber </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/david-frum/" class="c-menu__link" tabindex="-1"> David Frum </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="/author/adam-serwer/" class="c-menu__link" tabindex="-1"> Adam Serwer </a></li>  <li class="c-menu__item c-menu__item--quick-link"><a href="https://accounts.theatlantic.com/" class="c-menu__link" tabindex="-1"> Manage subscription </a></li>  </ul></section></div></div></div> </li></ul></div></nav><div class="c-overlay js-overlay is-hidden"></div>

        

        

        
<style type="text/css">

@media(min-width: 1592px) {
    .above-article,
    .photos .photo {
        width: 1500px;
    }
}




#img01 .img {
    padding-bottom: 65.4444444444%;
}


@media(min-width: 1592px) {
    #img01 .img {
        padding-bottom: 65.4%;
    }
}


@media(min-width: 1292px) {
    #img01 .img {
        padding-bottom: 65.4166666667%;
    }
}

@media(max-width: 992px) {
    #img01 .img {
        padding-bottom: 65.5%;
    }
}

@media(max-width: 692px) {
    #img01 .img {
        padding-bottom: 65.4444444444%;
    }
}

@media(max-width: 542px) {
    #img01 .img {
        padding-bottom: 65.5%;
    }
}





#img02 .img {
    padding-bottom: 63.5555555556%;
}


@media(min-width: 1592px) {
    #img02 .img {
        padding-bottom: 63.4666666667%;
    }
}


@media(min-width: 1292px) {
    #img02 .img {
        padding-bottom: 63.5%;
    }
}

@media(max-width: 992px) {
    #img02 .img {
        padding-bottom: 63.5%;
    }
}

@media(max-width: 692px) {
    #img02 .img {
        padding-bottom: 63.5555555556%;
    }
}

@media(max-width: 542px) {
    #img02 .img {
        padding-bottom: 63.5%;
    }
}





#img03 .img {
    padding-bottom: 67.5555555556%;
}


@media(min-width: 1592px) {
    #img03 .img {
        padding-bottom: 67.5333333333%;
    }
}


@media(min-width: 1292px) {
    #img03 .img {
        padding-bottom: 67.5%;
    }
}

@media(max-width: 992px) {
    #img03 .img {
        padding-bottom: 67.5%;
    }
}

@media(max-width: 692px) {
    #img03 .img {
        padding-bottom: 67.5555555556%;
    }
}

@media(max-width: 542px) {
    #img03 .img {
        padding-bottom: 67.5%;
    }
}





#img04 .img {
    padding-bottom: 71.5555555556%;
}


@media(min-width: 1592px) {
    #img04 .img {
        padding-bottom: 71.4666666667%;
    }
}


@media(min-width: 1292px) {
    #img04 .img {
        padding-bottom: 71.5%;
    }
}

@media(max-width: 992px) {
    #img04 .img {
        padding-bottom: 71.5%;
    }
}

@media(max-width: 692px) {
    #img04 .img {
        padding-bottom: 71.5555555556%;
    }
}

@media(max-width: 542px) {
    #img04 .img {
        padding-bottom: 71.5%;
    }
}





#img05 .img {
    padding-bottom: 71.6666666667%;
}


@media(min-width: 1592px) {
    #img05 .img {
        padding-bottom: 71.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img05 .img {
        padding-bottom: 71.6666666667%;
    }
}

@media(max-width: 992px) {
    #img05 .img {
        padding-bottom: 71.6666666667%;
    }
}

@media(max-width: 692px) {
    #img05 .img {
        padding-bottom: 71.6666666667%;
    }
}

@media(max-width: 542px) {
    #img05 .img {
        padding-bottom: 71.6666666667%;
    }
}





#img06 .img {
    padding-bottom: 58.2222222222%;
}


@media(min-width: 1592px) {
    #img06 .img {
        padding-bottom: 58.2%;
    }
}


@media(min-width: 1292px) {
    #img06 .img {
        padding-bottom: 58.1666666667%;
    }
}

@media(max-width: 992px) {
    #img06 .img {
        padding-bottom: 58.1666666667%;
    }
}

@media(max-width: 692px) {
    #img06 .img {
        padding-bottom: 58.2222222222%;
    }
}

@media(max-width: 542px) {
    #img06 .img {
        padding-bottom: 58.1666666667%;
    }
}





#img07 .img {
    padding-bottom: 75.0%;
}


@media(min-width: 1592px) {
    #img07 .img {
        padding-bottom: 75.0%;
    }
}


@media(min-width: 1292px) {
    #img07 .img {
        padding-bottom: 75.0%;
    }
}

@media(max-width: 992px) {
    #img07 .img {
        padding-bottom: 75.0%;
    }
}

@media(max-width: 692px) {
    #img07 .img {
        padding-bottom: 75.0%;
    }
}

@media(max-width: 542px) {
    #img07 .img {
        padding-bottom: 75.0%;
    }
}





#img08 .img {
    padding-bottom: 60.7777777778%;
}


@media(min-width: 1592px) {
    #img08 .img {
        padding-bottom: 60.7333333333%;
    }
}


@media(min-width: 1292px) {
    #img08 .img {
        padding-bottom: 60.75%;
    }
}

@media(max-width: 992px) {
    #img08 .img {
        padding-bottom: 60.8333333333%;
    }
}

@media(max-width: 692px) {
    #img08 .img {
        padding-bottom: 60.7777777778%;
    }
}

@media(max-width: 542px) {
    #img08 .img {
        padding-bottom: 60.8333333333%;
    }
}





#img09 .img {
    padding-bottom: 62.4444444444%;
}


@media(min-width: 1592px) {
    #img09 .img {
        padding-bottom: 62.4%;
    }
}


@media(min-width: 1292px) {
    #img09 .img {
        padding-bottom: 62.4166666667%;
    }
}

@media(max-width: 992px) {
    #img09 .img {
        padding-bottom: 62.5%;
    }
}

@media(max-width: 692px) {
    #img09 .img {
        padding-bottom: 62.4444444444%;
    }
}

@media(max-width: 542px) {
    #img09 .img {
        padding-bottom: 62.5%;
    }
}





#img10 .img {
    padding-bottom: 66.3333333333%;
}


@media(min-width: 1592px) {
    #img10 .img {
        padding-bottom: 66.3333333333%;
    }
}


@media(min-width: 1292px) {
    #img10 .img {
        padding-bottom: 66.3333333333%;
    }
}

@media(max-width: 992px) {
    #img10 .img {
        padding-bottom: 66.3333333333%;
    }
}

@media(max-width: 692px) {
    #img10 .img {
        padding-bottom: 66.3333333333%;
    }
}

@media(max-width: 542px) {
    #img10 .img {
        padding-bottom: 66.3333333333%;
    }
}





#img11 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img11 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img11 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img11 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img11 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img11 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img12 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img12 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img12 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img12 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img12 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img12 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img13 .img {
    padding-bottom: 65.0%;
}


@media(min-width: 1592px) {
    #img13 .img {
        padding-bottom: 65.0%;
    }
}


@media(min-width: 1292px) {
    #img13 .img {
        padding-bottom: 65.0%;
    }
}

@media(max-width: 992px) {
    #img13 .img {
        padding-bottom: 65.0%;
    }
}

@media(max-width: 692px) {
    #img13 .img {
        padding-bottom: 65.0%;
    }
}

@media(max-width: 542px) {
    #img13 .img {
        padding-bottom: 65.0%;
    }
}





#img14 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img14 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img14 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img14 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img14 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img14 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img15 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img15 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img15 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img15 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img15 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img15 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img16 .img {
    padding-bottom: 67.2222222222%;
}


@media(min-width: 1592px) {
    #img16 .img {
        padding-bottom: 67.1333333333%;
    }
}


@media(min-width: 1292px) {
    #img16 .img {
        padding-bottom: 67.1666666667%;
    }
}

@media(max-width: 992px) {
    #img16 .img {
        padding-bottom: 67.1666666667%;
    }
}

@media(max-width: 692px) {
    #img16 .img {
        padding-bottom: 67.2222222222%;
    }
}

@media(max-width: 542px) {
    #img16 .img {
        padding-bottom: 67.1666666667%;
    }
}





#img17 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img17 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img17 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img17 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img17 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img17 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img18 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img18 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img18 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img18 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img18 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img18 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img19 .img {
    padding-bottom: 88.2222222222%;
}


@media(min-width: 1592px) {
    #img19 .img {
        padding-bottom: 88.2666666667%;
    }
}


@media(min-width: 1292px) {
    #img19 .img {
        padding-bottom: 88.25%;
    }
}

@media(max-width: 992px) {
    #img19 .img {
        padding-bottom: 88.1666666667%;
    }
}

@media(max-width: 692px) {
    #img19 .img {
        padding-bottom: 88.2222222222%;
    }
}

@media(max-width: 542px) {
    #img19 .img {
        padding-bottom: 88.1666666667%;
    }
}





#img20 .img {
    padding-bottom: 61.5555555556%;
}


@media(min-width: 1592px) {
    #img20 .img {
        padding-bottom: 61.5333333333%;
    }
}


@media(min-width: 1292px) {
    #img20 .img {
        padding-bottom: 61.5%;
    }
}

@media(max-width: 992px) {
    #img20 .img {
        padding-bottom: 61.5%;
    }
}

@media(max-width: 692px) {
    #img20 .img {
        padding-bottom: 61.5555555556%;
    }
}

@media(max-width: 542px) {
    #img20 .img {
        padding-bottom: 61.5%;
    }
}





#img21 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img21 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img21 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img21 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img21 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img21 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img22 .img {
    padding-bottom: 68.7777777778%;
}


@media(min-width: 1592px) {
    #img22 .img {
        padding-bottom: 68.7333333333%;
    }
}


@media(min-width: 1292px) {
    #img22 .img {
        padding-bottom: 68.75%;
    }
}

@media(max-width: 992px) {
    #img22 .img {
        padding-bottom: 68.8333333333%;
    }
}

@media(max-width: 692px) {
    #img22 .img {
        padding-bottom: 68.7777777778%;
    }
}

@media(max-width: 542px) {
    #img22 .img {
        padding-bottom: 68.8333333333%;
    }
}





#img23 .img {
    padding-bottom: 66.7777777778%;
}


@media(min-width: 1592px) {
    #img23 .img {
        padding-bottom: 66.7333333333%;
    }
}


@media(min-width: 1292px) {
    #img23 .img {
        padding-bottom: 66.75%;
    }
}

@media(max-width: 992px) {
    #img23 .img {
        padding-bottom: 66.8333333333%;
    }
}

@media(max-width: 692px) {
    #img23 .img {
        padding-bottom: 66.7777777778%;
    }
}

@media(max-width: 542px) {
    #img23 .img {
        padding-bottom: 66.8333333333%;
    }
}





#img24 .img {
    padding-bottom: 63.4444444444%;
}


@media(min-width: 1592px) {
    #img24 .img {
        padding-bottom: 63.4%;
    }
}


@media(min-width: 1292px) {
    #img24 .img {
        padding-bottom: 63.4166666667%;
    }
}

@media(max-width: 992px) {
    #img24 .img {
        padding-bottom: 63.5%;
    }
}

@media(max-width: 692px) {
    #img24 .img {
        padding-bottom: 63.4444444444%;
    }
}

@media(max-width: 542px) {
    #img24 .img {
        padding-bottom: 63.5%;
    }
}





#img25 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img25 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img25 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img25 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img25 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img25 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img26 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img26 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img26 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img26 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img26 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img26 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img27 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img27 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img27 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img27 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img27 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img27 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img28 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img28 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img28 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img28 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img28 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img28 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img29 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img29 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img29 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img29 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img29 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img29 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img30 .img {
    padding-bottom: 68.7777777778%;
}


@media(min-width: 1592px) {
    #img30 .img {
        padding-bottom: 68.7333333333%;
    }
}


@media(min-width: 1292px) {
    #img30 .img {
        padding-bottom: 68.75%;
    }
}

@media(max-width: 992px) {
    #img30 .img {
        padding-bottom: 68.8333333333%;
    }
}

@media(max-width: 692px) {
    #img30 .img {
        padding-bottom: 68.7777777778%;
    }
}

@media(max-width: 542px) {
    #img30 .img {
        padding-bottom: 68.8333333333%;
    }
}





#img31 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img31 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img31 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img31 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img31 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img31 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img32 .img {
    padding-bottom: 66.7777777778%;
}


@media(min-width: 1592px) {
    #img32 .img {
        padding-bottom: 66.7333333333%;
    }
}


@media(min-width: 1292px) {
    #img32 .img {
        padding-bottom: 66.75%;
    }
}

@media(max-width: 992px) {
    #img32 .img {
        padding-bottom: 66.8333333333%;
    }
}

@media(max-width: 692px) {
    #img32 .img {
        padding-bottom: 66.7777777778%;
    }
}

@media(max-width: 542px) {
    #img32 .img {
        padding-bottom: 66.8333333333%;
    }
}





#img33 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img33 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img33 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img33 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img33 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img33 .img {
        padding-bottom: 66.6666666667%;
    }
}





#img34 .img {
    padding-bottom: 66.7777777778%;
}


@media(min-width: 1592px) {
    #img34 .img {
        padding-bottom: 66.7333333333%;
    }
}


@media(min-width: 1292px) {
    #img34 .img {
        padding-bottom: 66.75%;
    }
}

@media(max-width: 992px) {
    #img34 .img {
        padding-bottom: 66.8333333333%;
    }
}

@media(max-width: 692px) {
    #img34 .img {
        padding-bottom: 66.7777777778%;
    }
}

@media(max-width: 542px) {
    #img34 .img {
        padding-bottom: 66.8333333333%;
    }
}





#img35 .img {
    padding-bottom: 66.6666666667%;
}


@media(min-width: 1592px) {
    #img35 .img {
        padding-bottom: 66.6666666667%;
    }
}


@media(min-width: 1292px) {
    #img35 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 992px) {
    #img35 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 692px) {
    #img35 .img {
        padding-bottom: 66.6666666667%;
    }
}

@media(max-width: 542px) {
    #img35 .img {
        padding-bottom: 66.6666666667%;
    }
}




</style>

<main class="article-container" id="article">

<div class="ad-wrapper ad-stickyleader-wrapper fix-available">
    <gpt-ad
        targeting-pos="stickyleader-photo"
        class="ad ad-stickyleader-photo"
        prevent-reload="1"
        >
        <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250]]"></gpt-sizeset>
        <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90]]"></gpt-sizeset>
    </gpt-ad>
</div>

<div class="above-article" id="main-content">
    <div id="article-header">
        <h1 class="hed">Photos of the Week: Frozen Road, Broadway Gorilla, Penguin Swing Set</h1>
        <ul class="metadata">
            <li class="byline"><a href="/author/alan-taylor/" title="Alan Taylor">Alan Taylor</a></li>
            <li class="date">Mar 8, 2019</li>
            <li class="num-photos">35 Photos</li>
            <li class="kicker">In Focus</li>
        </ul>

        <div id="article-content">
            <p>Power outages in Venezuela, Mardi Gras festivities in New Orleans, skijoring in Colorado, Carnival celebrations in Brazil, the Crufts dog show in England, snowboarding in California, presidential campaigning in New York, baseball spring training in Arizona, International Women&#8217;s Day observed around the world, and much more</p>
        </div>
        <a href="#" data-omni-click="r'photo',r' ',d,r'read more',r' ',l.pathname" class="read-more">Read more</a>

        <div class="hints">
            <strong>Hints:</strong>
            <span class="detect-fullscreen">View this page <a href="#" class="fullscreen">full screen</a>.</span>
            
                Skip to the next and previous photo by typing j/k or &larr;/&rarr;.
            
        </div>
        

<ul class="social-icons round color component-social ">
    <li><a href="#" class="social-icon facebook" data-share="facebook" data-omni-click="r'articlefb',r'',l.pathname"><span>Facebook</span></a></li>
    <li><a href="#" class="social-icon twitter" data-share="twitter" data-share-via="TheAtlantic" data-omni-click="r'articletweet',r'',l.pathname"><span>Twitter</span></a></li>
    <li><a href="#" class="social-icon linkedin" data-share="linkedin" data-omni-click="r'articleli',r'',l.pathname"><span>LinkedIn</span></a></li>
    <li><a href="mailto:?subject=The%20Atlantic: Photos%20of%20the%20Week%3A%20Frozen%20Road%2C%20Broadway%20Gorilla%2C%20Penguin%20Swing%20Set
                &body=Power%20outages%20in%20Venezuela%2C%20Mardi%20Gras%20festivities%20in%20New%20Orleans%2C%20presidential%20campaigning%20in%20New%20York%2C%20International%20Women%E2%80%99s%20Day%20observed%20around%20the%20world%2C%20and%20much%20more
                %0A%0A
                Read this:%0A
                https://www.theatlantic.com/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/%3Futm_source%3Datl%26utm_medium%3Demail%26utm_campaign%3Dshare" class="social-icon email" data-omni-click="r'articleemail',r'',l.pathname" target="_blank"><span>Email/span></a></li>
</ul>

    </div>
</div>

<ul class="photos" data-expand="2000"> 
    
    
    <li class="photo"
            id="img01">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w01_1133158845/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Whitney Barron greets Chrystall Cooper during the annual Mardi Gras parade in Sydney, Australia, on March 2, 2019. It is the first time the charity Haka for Life has had a float in the parade. The float is designed like a Waka canoe, which is significant in Maori culture, and has the theme of "One Waka, One Love, We&#x2019;re All in the Same Boat Together." Haka for Life is an organization focused on men's health, well-being, and suicide prevention in the Maori community. The organization believes that there is a stigma in Maori warrior culture that implies that people who identify as LGBTQ are not warriors. The Sydney Mardi Gras parade began in 1978 as a commemoration of the 1969 Stonewall riots in New York. The annual event promotes the awareness of gay, lesbian, bisexual, and transgender issues and themes.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img01">#</a>
            </p>
            <div class="credit">
                
                
                Lisa Maree Williams / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    <li class="ad-wrapper">
        
        <gpt-ad
        targeting-pos="bigad1"
        class="ad ad-bigad"
        >
        
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[728, 90], [970, 250], [1, 3]]"></gpt-sizeset>
        </gpt-ad>
    </li>
    
    
    <li class="photo
            
            
            "
            id="img02">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w02_1129218742/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w02_1129218742/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w02_1129218742/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w02_1129218742/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w02_1129218742/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w02_1129218742/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>A feminist activist shouts into a microphone during a rally for gender equality and women's rights in St. Petersburg, Russia, on March 8, 2019, as International Women's Day is celebrated around the world.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img02">#</a>
            </p>
            <div class="credit">
                
                
                Olga Maltseva / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img03">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w03_1128958183/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w03_1128958183/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w03_1128958183/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w03_1128958183/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w03_1128958183/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w03_1128958183/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Kong and cast members bow during curtain call as the mayor's office honors <i>King Kong</i> with a mayoral message and ceremonial street renaming at the Broadway Theatre in New York City on March 6, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img03">#</a>
            </p>
            <div class="credit">
                
                
                Angela Weiss / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img04">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w04_1128496602/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w04_1128496602/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w04_1128496602/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w04_1128496602/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w04_1128496602/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w04_1128496602/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>A child dressed as the pope cries during the weekly Angelus prayer at St. Peter's Square at the Vatican on March 3, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img04">#</a>
            </p>
            <div class="credit">
                
                
                Tiziana Fabi / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img05">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w05_AP19067007357822/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w05_AP19067007357822/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w05_AP19067007357822/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w05_AP19067007357822/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w05_AP19067007357822/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w05_AP19067007357822/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The 6-year-old Joan Stiltner of Chicago helps hold a ball for the Chicago White Sox left fielder Joel Booker to autograph before the team's spring-training baseball game against the Milwaukee Brewers in Glendale, Arizona, on March 7, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img05">#</a>
            </p>
            <div class="credit">
                
                
                Sue Ogrocki / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img06">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w06_AP19067463932627/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w06_AP19067463932627/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w06_AP19067463932627/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w06_AP19067463932627/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w06_AP19067463932627/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w06_AP19067463932627/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>An artwork by the British artist Damien Hirst titled <i>The Incomplete Truth</i>&#x2014;made of glass, painted aluminum, silicone, acrylic, stainless steel, dove, and formaldehyde solution, and owned by the late British vocalist George Michael&#x2014;is displayed during a press preview at Christie's auction house in London on March 8, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img06">#</a>
            </p>
            <div class="credit">
                
                
                Alastair Grant / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    <li class="ad-wrapper">
        
        <gpt-ad
        targeting-pos="bigad2"
        class="ad ad-bigad"
        >
        
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[728, 90], [970, 250], [1, 3]]"></gpt-sizeset>
        </gpt-ad>
    </li>
    
    
    <li class="photo
            
            expanded
            "
            id="img07">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w07_1128822211/main_1500.jpg?1552072075" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w07_1128822211/main_1200.jpg?1552072075" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w07_1128822211/main_600.jpg?1552072075" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w07_1128822211/main_900.jpg?1552072075" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w07_1128822211/main_600.jpg?1552072075" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w07_1128822211/main_900.jpg?1552072075" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>A rainwater reservoir overflows into a fan-shaped waterfall in Guilin, Guangxi, China, on March 5, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img07">#</a>
            </p>
            <div class="credit">
                
                
                Barcroft Media via Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img08">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w08_RTX6Q64X/main_1500.jpg?1552077539" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w08_RTX6Q64X/main_1200.jpg?1552077539" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w08_RTX6Q64X/main_600.jpg?1552077539" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w08_RTX6Q64X/main_900.jpg?1552077539" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w08_RTX6Q64X/main_600.jpg?1552077539" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w08_RTX6Q64X/main_900.jpg?1552077539" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>A car drives along a five-kilometer-long road, which is supervised by regional traffic services and connects the banks of the ice-covered Yenisei River during the winter season, south of Krasnoyarsk, Russia, on March 7, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img08">#</a>
            </p>
            <div class="credit">
                
                
                Ilya Naymushin / Reuters
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img09">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w09_1128412471/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w09_1128412471/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w09_1128412471/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w09_1128412471/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w09_1128412471/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w09_1128412471/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The eagle mascot of the Lazio soccer team flies over the Olympic Stadium in Rome prior to the Italian Serie A soccer match between S.S. Lazio and A.S. Roma on March 2, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img09">#</a>
            </p>
            <div class="credit">
                
                
                Tiziana Fabi / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img10">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w10_AP19065528743676/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w10_AP19065528743676/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w10_AP19065528743676/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w10_AP19065528743676/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w10_AP19065528743676/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w10_AP19065528743676/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>A Balinese man reacts as he is hit with flaming coconut leaves during the firefight ritual called "Lukat Gni" before Nyepi in Klungkung, Bali, Indonesia, on March 6, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img10">#</a>
            </p>
            <div class="credit">
                
                
                Firdia Lisnawati / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img11">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w11_1133239401/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w11_1133239401/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w11_1133239401/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w11_1133239401/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w11_1133239401/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w11_1133239401/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The Democratic presidential candidate Senator Bernie Sanders holds his first presidential-campaign rally at Brooklyn College in New York on March 2, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img11">#</a>
            </p>
            <div class="credit">
                
                
                Kena Betancur / VIEWpress / Corbis via Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    <li class="ad-wrapper">
        
        <gpt-ad
        targeting-pos="bigad3"
        class="ad ad-bigad"
        >
        
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[728, 90], [970, 250], [1, 3]]"></gpt-sizeset>
        </gpt-ad>
    </li>
    
    
    <li class="photo
            
            expanded
            "
            id="img12">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w12_AP19061637819666/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w12_AP19061637819666/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w12_AP19061637819666/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w12_AP19061637819666/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w12_AP19061637819666/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w12_AP19061637819666/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>President Donald Trump hugs the American flag as he arrives to speak at the Conservative Political Action Conference in Oxon Hill, Maryland, on March 2, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img12">#</a>
            </p>
            <div class="credit">
                
                
                Carolyn Kaster / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img13">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w13_AP19063663588594/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w13_AP19063663588594/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w13_AP19063663588594/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w13_AP19063663588594/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w13_AP19063663588594/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w13_AP19063663588594/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>President Donald Trump welcomes the 2018 NCAA FCS college-football champions, the North Dakota State Bison, to the State Dining Room of the White House in Washington, D.C., with McDonald's and Chick-fil-A fast food on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img13">#</a>
            </p>
            <div class="credit">
                
                
                Carolyn Kaster / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img14">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w14_RTX6Q6SV/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w14_RTX6Q6SV/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w14_RTX6Q6SV/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w14_RTX6Q6SV/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w14_RTX6Q6SV/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w14_RTX6Q6SV/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>A group of Central American migrants prepares to surrender to U.S. Border Patrol agents south of the U.S.-Mexico border fence in El Paso, Texas, on March 6, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img14">#</a>
            </p>
            <div class="credit">
                
                
                Lucy Nicholson / Reuters
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img15">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w15_1128552130/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w15_1128552130/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w15_1128552130/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w15_1128552130/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w15_1128552130/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w15_1128552130/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The runway is pictured at the end of the Givenchy fall 2019/winter 2020 ready-to-wear fashion show in Paris on March 3, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img15">#</a>
            </p>
            <div class="credit">
                
                
                Christophe Archambault / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img16">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w16_1134374658/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w16_1134374658/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w16_1134374658/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w16_1134374658/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w16_1134374658/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w16_1134374658/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Yuto Totsuka of Japan takes a practice run before the start of the snowboard-halfpipe qualifiers at the 2019 U.S. Grand Prix at Mammoth Mountain in California on March 7, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img16">#</a>
            </p>
            <div class="credit">
                
                
                Ezra Shaw / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    <li class="ad-wrapper">
        
        <gpt-ad
        targeting-pos="bigad4"
        class="ad ad-bigad"
        >
        
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[728, 90], [970, 250], [1, 3]]"></gpt-sizeset>
        </gpt-ad>
    </li>
    
    
    <li class="photo
            
            
            "
            id="img17">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w17_1133373496/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w17_1133373496/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w17_1133373496/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w17_1133373496/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w17_1133373496/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w17_1133373496/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Ana Peleteiro of Spain is photographed on her way to winning gold in the women's triple-jump final during day three of the 2019 European Athletics Indoor Championships at Emirates Arena in Glasgow, Scotland, on March 3, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img17">#</a>
            </p>
            <div class="credit">
                
                
                Michael Steele / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img18">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w18_1133176879/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w18_1133176879/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w18_1133176879/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w18_1133176879/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w18_1133176879/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w18_1133176879/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Freshmen at Zhejiang GongShang University receive military training in Hangzhou, Zhejiang province, China, on March 2, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img18">#</a>
            </p>
            <div class="credit">
                
                
                Xu Junyong / VCG via Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img19">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w19_AP19066828597638/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w19_AP19066828597638/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w19_AP19066828597638/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w19_AP19066828597638/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w19_AP19066828597638/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w19_AP19066828597638/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Lucas Romero <i>right</i> of Brazil's Cruzeiro fights for the ball with Iv&#xE1;n Javier Rossi of Argentina's Hurac&#xE1;n in a downpour during a Copa Libertadores soccer game in Buenos Aires, Argentina, on March 7, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img19">#</a>
            </p>
            <div class="credit">
                
                
                Natacha Pisarenko / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img20">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w20_1128607255/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w20_1128607255/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w20_1128607255/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w20_1128607255/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w20_1128607255/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w20_1128607255/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Revelers dance as they take part in the "Los Indianos" street carnival in Santa Cruz de la Palma, on the Spanish Canary Island of La Palma, on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img20">#</a>
            </p>
            <div class="credit">
                
                
                Desiree Martin / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img21">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w21_1128524889/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w21_1128524889/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w21_1128524889/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w21_1128524889/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w21_1128524889/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w21_1128524889/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Waves crash against a harbor wall during Storm Freya in Porthcawl, Wales, on March 3, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img21">#</a>
            </p>
            <div class="credit">
                
                
                Matthew Horwood / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    <li class="ad-wrapper">
        
        <gpt-ad
        targeting-pos="bigad5"
        class="ad ad-bigad"
        >
        
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[728, 90], [970, 250], [1, 3]]"></gpt-sizeset>
        </gpt-ad>
    </li>
    
    
    <li class="photo
            
            expanded
            "
            id="img22">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w22_RTX6Q7WB/main_1500.jpg?1552079798" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w22_RTX6Q7WB/main_1200.jpg?1552079798" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w22_RTX6Q7WB/main_600.jpg?1552079798" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w22_RTX6Q7WB/main_900.jpg?1552079798" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w22_RTX6Q7WB/main_600.jpg?1552079798" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w22_RTX6Q7WB/main_900.jpg?1552079798" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>An Old English sheepdog is groomed during the second day of the Crufts dog show in Birmingham, England, on March 8, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img22">#</a>
            </p>
            <div class="credit">
                
                
                Hannah Mckay / Reuters
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img23">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w23_AP19064107106415/main_1500.jpg?1552072075" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w23_AP19064107106415/main_1200.jpg?1552072075" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w23_AP19064107106415/main_600.jpg?1552072075" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w23_AP19064107106415/main_900.jpg?1552072075" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w23_AP19064107106415/main_600.jpg?1552072075" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w23_AP19064107106415/main_900.jpg?1552072075" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Performers from the Vila Isabel samba school parade during Carnival celebrations at the Sambadrome in Rio de Janeiro, Brazil, on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img23">#</a>
            </p>
            <div class="credit">
                
                
                Silvia Izquierdo / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img24">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w24_AP19063212927469/main_1500.jpg?1552072075" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w24_AP19063212927469/main_1200.jpg?1552072075" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w24_AP19063212927469/main_600.jpg?1552072075" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w24_AP19063212927469/main_900.jpg?1552072075" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w24_AP19063212927469/main_600.jpg?1552072075" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w24_AP19063212927469/main_900.jpg?1552072075" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Performers from the Salgueiro samba school parade during Carnival at the Sambadrome in Rio on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img24">#</a>
            </p>
            <div class="credit">
                
                
                Silvia Izquierdo / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img25">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w25_RTX6POQC/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w25_RTX6POQC/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w25_RTX6POQC/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w25_RTX6POQC/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w25_RTX6POQC/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w25_RTX6POQC/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Revelers from the Vila Isabel samba school perform during the second night of the Carnival parade at the Sambadrome in Rio on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img25">#</a>
            </p>
            <div class="credit">
                
                
                Pilar Olivares / Reuters
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img26">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w26_1134278729/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w26_1134278729/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w26_1134278729/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w26_1134278729/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w26_1134278729/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w26_1134278729/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The jockey Harry Kimber is momentarily trapped under his mount, West Chinnock, after a fall (both the horse and jockey were okay) at the last hurdle at Wincanton Racecourse in England on March 7, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img26">#</a>
            </p>
            <div class="credit">
                
                
                Alan Crowhurst / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    <li class="ad-wrapper">
        
        <gpt-ad
        targeting-pos="bigad6"
        class="ad ad-bigad"
        >
        
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[728, 90], [970, 250], [1, 3]]"></gpt-sizeset>
        </gpt-ad>
    </li>
    
    
    <li class="photo
            
            
            "
            id="img27">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w27_1128438973/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w27_1128438973/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w27_1128438973/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w27_1128438973/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w27_1128438973/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w27_1128438973/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Jeff Dahl races down Harrison Avenue while the skier Vincent Pestello loses control in the first jump of the skijoring course at the 71st annual Leadville Ski Joring weekend competition in Leadville, Colorado, on March 2, 2019. Skijoring, which has its origins as a competitive sport in Scandinavia, has been adapted over the years to include jumps, slalom gates, and spearing rings for points. Leadville has been hosting skijoring competitions since 1949.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img27">#</a>
            </p>
            <div class="credit">
                
                
                Jason Connolly / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img28">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w28_AP19067396740791/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w28_AP19067396740791/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w28_AP19067396740791/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w28_AP19067396740791/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w28_AP19067396740791/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w28_AP19067396740791/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>A journalist live-streams using multiple devices outside the Great Hall of the People, where a session of the National People's Congress was being held, in Beijing, China, on March 8, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img28">#</a>
            </p>
            <div class="credit">
                
                
                Ng Han Guan / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img29">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w29_AP19062383247333/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w29_AP19062383247333/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w29_AP19062383247333/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w29_AP19062383247333/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w29_AP19062383247333/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w29_AP19062383247333/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The former NBA basketball player Yao Ming, a delegate to the Chinese People's Political Consultative Conference (CPPCC), towers over other delegates as they arrive at the Great Hall of the People to attend the opening session of the CPPCC in Beijing on March 3, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img29">#</a>
            </p>
            <div class="credit">
                
                
                Andy Wong / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img30">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w30_RTX6PLHE/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w30_RTX6PLHE/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w30_RTX6PLHE/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w30_RTX6PLHE/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w30_RTX6PLHE/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w30_RTX6PLHE/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>People dressed as penguins play at a playground on Rosenmontag, or "Rose Monday," in Cologne, Germany, on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img30">#</a>
            </p>
            <div class="credit">
                
                
                Thilo Schmuelgen / Reuters
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img31">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w31_1133974603/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w31_1133974603/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w31_1133974603/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w31_1133974603/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w31_1133974603/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w31_1133974603/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Members of the Black Mohawks Mardi Gras Indians "mask" in New Orleans, Louisiana, on March 5, 2019. Part of Mardi Gras tradition in New Orleans involves dozens of neighborhood tribes of <a href="https://en.wikipedia.org/wiki/Mardi_Gras_Indians">Mardi Gras Indians</a> "masking," or parading, in the streets.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img31">#</a>
            </p>
            <div class="credit">
                
                
                Erika Goldring / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    <li class="ad-wrapper">
        
        <gpt-ad
        targeting-pos="bigad7"
        class="ad ad-bigad"
        >
        
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[760, 0]" sizes="[[728, 90], [1, 3]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[728, 90], [970, 250], [1, 3]]"></gpt-sizeset>
        </gpt-ad>
    </li>
    
    
    <li class="photo
            
            expanded
            "
            id="img32">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w32_1128642575/main_1500.jpg?1552072075" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w32_1128642575/main_1200.jpg?1552072075" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w32_1128642575/main_600.jpg?1552072075" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w32_1128642575/main_900.jpg?1552072075" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w32_1128642575/main_600.jpg?1552072075" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w32_1128642575/main_900.jpg?1552072075" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>Revelers take part in the annual clown parade in Sesimbra, Portugal, on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img32">#</a>
            </p>
            <div class="credit">
                
                
                Patricia De Melo Moreira / AFP / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img33">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w33_1133763068/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w33_1133763068/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w33_1133763068/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w33_1133763068/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w33_1133763068/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w33_1133763068/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The Orpheus Leviathan float, with smoke and fiber-optic lights, rolls down Napoleon Avenue in the 2019 Krewe of Orpheus parade in New Orleans, Louisiana, on March 4, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img33">#</a>
            </p>
            <div class="credit">
                
                
                Erika Goldring / Getty
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            
            "
            id="img34">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w34_AP19067070797032/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w34_AP19067070797032/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w34_AP19067070797032/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w34_AP19067070797032/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w34_AP19067070797032/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w34_AP19067070797032/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>In Caracas, Venezuela, relatives of a patient walk in the darkened hall of a clinic with candles lighting the way during a power outage on March 7, 2019. A power outage left much of Venezuela in the dark early Thursday evening, in what appeared to be one of the largest blackouts yet in a country where power failures have become increasingly common.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img34">#</a>
            </p>
            <div class="credit">
                
                
                Ariana Cubillos / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <li class="photo
            
            expanded
            "
            id="img35">

        <figure class="img">
            <picture>
                
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w35_AP19066038954319/main_1500.jpg?1552077378" media="(min-width: 1592px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w35_AP19066038954319/main_1200.jpg?1552077378" media="(min-width: 1292px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w35_AP19066038954319/main_600.jpg?1552077378" media="(max-width: 542px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w35_AP19066038954319/main_900.jpg?1552077378" media="(max-width: 692px)">
                <source data-srcset="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w35_AP19066038954319/main_600.jpg?1552077378" media="(max-width: 992px)">
                <img data-src="https://cdn.theatlantic.com/assets/media/img/photo/2019/03/photos-of-the-week/w35_AP19066038954319/main_900.jpg?1552077378" width="100%" class="lazyload">
            </picture>

            

            
        </figure>

        <div class="mobile-caption-wrapper">
            <p class="caption">
                <span>The sun rises over the Penang Bridge on Penang Island, northern Malaysia, on March 7, 2019.</span>
                <a class="permalink" href="/photo/2019/03/photos-of-the-week-frozen-road-broadway-gorilla-penguin-swingset/584431/#img35">#</a>
            </p>
            <div class="credit">
                
                
                Vincent Thian / AP
                
            </div>
        </div>
        <a href="#" class="read-more">Read more</a>
        
        
    </li>

    
    
    <section class="c-letters-cta"><p class="c-letters-cta__text">We want to hear what you think about this article. <a href="https://www.theatlantic.com/contact/letters/" class="c-letters-cta__link">Submit a letter</a> to the editor or write to letters@theatlantic.com.</p></section>
</ul>


<script>
    $('.transition-img').click(function () {
        $(this).toggleClass('active');
    });
</script>

</main>







<div class="fluid-container" id="below-article">
    <div class="below-photo-article-ad-wrapper">
        <div class="sponsored-links-container">
            <gpt-ad targeting-pos="sponsored-links"  lazy-load="0"><gpt-sizeset viewport-size="[0, 0]" sizes="sponsored-links"></gpt-sizeset></gpt-ad>
        </div>
        <gpt-ad
            targeting-pos="midboxrightphoto"
            class="ad ad-midboxrightphoto"
            >
            <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250]]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[680, 0]" sizes="[]"></gpt-sizeset>
            <gpt-sizeset viewport-size="[1010, 0]" sizes="[[300, 250]]"></gpt-sizeset>
        </gpt-ad>
    </div>

    <div id="mid-promo-wrapper">
            <div id="promo" data-omni-click="r'photo',$li,d,r'most-pop',r' ',l.pathname">
                
                    

<div class="module module-recent">
    <h2 class="module-label">Most Recent</h2>
    <div class="grid-wrapper">
        <ul class="grid">
            
            <li id="recent1" class="grid-article recent-article two-row">
                <style type="text/css">
                    #recent1 .lead-image {
                        
                        background-image: url(https://cdn.theatlantic.com/assets/media/img/photo/2019/03/sri-lanka-tk/s01_1129004497-1/river_thumb.jpg?1551983205);
                    }
                </style>
                
<a href="/photo/2019/03/the-women-who-are-clearing-the-minefields-in-sri-lanka/584368/" data-omni-click="inherit">
    <figure>
        <div class="lead-image" title=""></div>
        <figcaption class="credit">Allison Joyce / Getty</figcaption>
    </figure>
</a>

<ul class="metadata">
    <li class="kicker">In Focus</li>
    <li class="date">Mar 7, 2019</li>
    <li class="num-photos">21 Photos</li>
</ul>

<h1 class="hed">
    <a data-omni-click="inherit" href="/photo/2019/03/the-women-who-are-clearing-the-minefields-in-sri-lanka/584368/">
        The Women Who Are Clearing the Minefields in Sri Lanka
    </a>
</h1>
<p class="dek">Allison Joyce, a photographer with Getty Images, spent time with some of the women taking on the dangerous but necessary task of clearing Sri Lanka’s remaining minefields.</p>

            </li>
            
            <li id="recent2" class="grid-article recent-article">
                <style type="text/css">
                    #recent2 .lead-image {
                        
                        background-image: url(https://cdn.theatlantic.com/assets/media/img/photo/2019/03/alabama-tornadoes-tk/a01_AP19063620921300-1/river_thumb.jpg?1551895713);
                    }
                </style>
                
<a href="/photo/2019/03/alabama-tornadoes-photos-aftermath/584260/" data-omni-click="inherit">
    <figure>
        <div class="lead-image" title=""></div>
        <figcaption class="credit">David Goldman / AP</figcaption>
    </figure>
</a>

<ul class="metadata">
    <li class="kicker">In Focus</li>
    <li class="date">Mar 6, 2019</li>
    <li class="num-photos">28 Photos</li>
</ul>

<h1 class="hed">
    <a data-omni-click="inherit" href="/photo/2019/03/alabama-tornadoes-photos-aftermath/584260/">
        Alabama Tornado Devastation in Photos
    </a>
</h1>
<p class="dek">On March 3, a monster of a tornado left a path of destruction more than 20 miles long in Lee County, Alabama.</p>

            </li>
            
            <li id="recent3" class="grid-article recent-article two-row">
                <style type="text/css">
                    #recent3 .lead-image {
                        
                        background-image: url(https://cdn.theatlantic.com/assets/media/img/photo/2019/03/finalists-smithsonian-magazines-201/s01_AE_RoryDoyle-1/river_thumb.jpg?1551808077);
                    }
                </style>
                
<a href="/photo/2019/03/finalists-smithsonian-magazines-2018-photo-contest/584182/" data-omni-click="inherit">
    <figure>
        <div class="lead-image" title=""></div>
        <figcaption class="credit">© Rory Doyle. All rights reserved.</figcaption>
    </figure>
</a>

<ul class="metadata">
    <li class="kicker">In Focus</li>
    <li class="date">Mar 5, 2019</li>
    <li class="num-photos">15 Photos</li>
</ul>

<h1 class="hed">
    <a data-omni-click="inherit" href="/photo/2019/03/finalists-smithsonian-magazines-2018-photo-contest/584182/">
        Finalists From <em>Smithsonian</em> Magazine’s 2018 Photo Contest
    </a>
</h1>
<p class="dek"><em>Smithsonian </em>magazine just announced the group of 60 finalists in its 16th annual photo contest.</p>

            </li>
            
            <li id="recent4" class="grid-article recent-article">
                <style type="text/css">
                    #recent4 .lead-image {
                        
                        background-image: url(https://cdn.theatlantic.com/assets/media/img/photo/2019/03/brazil-carnival-2019-photos/c01_1128576841-1/river_thumb.jpg?1551725608);
                    }
                </style>
                
<a href="/photo/2019/03/brazil-carnival-2019-photos/584050/" data-omni-click="inherit">
    <figure>
        <div class="lead-image" title=""></div>
        <figcaption class="credit">Buda Mendes / Getty</figcaption>
    </figure>
</a>

<ul class="metadata">
    <li class="kicker">In Focus</li>
    <li class="date">Mar 4, 2019</li>
    <li class="num-photos">30 Photos</li>
</ul>

<h1 class="hed">
    <a data-omni-click="inherit" href="/photo/2019/03/brazil-carnival-2019-photos/584050/">
        Carnival 2019 in Brazil
    </a>
</h1>
<p class="dek">Colorful images from Rio, São Paulo, and more</p>

            </li>
            
        </ul>
    </div>
</div>

                
                
                    
<section class="most-popular module module-most-popular">
    <h2 class="module-label">Most Popular on The Atlantic</h2>

    <ol>
        
        <li><a href="https://www.theatlantic.com/family/archive/2019/03/eating-the-same-thing-lunch-meal/584347/" data-omni-click="inherit">The People Who Eat the Same Meal Every Day </a></li>
        
        <li><a href="https://www.theatlantic.com/entertainment/archive/2019/03/leaving-neverland-how-michael-jackson-used-celebrity/584089/" data-omni-click="inherit">On Not Believing<em> Leaving Neverland</em> </a></li>
        
        <li><a href="https://www.theatlantic.com/health/archive/2019/03/typhus-tuberculosis-medieval-diseases-spreading-homeless/584380/" data-omni-click="inherit">Medieval Diseases Are Infecting California’s Homeless </a></li>
        
        <li><a href="https://www.theatlantic.com/education/archive/2019/03/decline-men-without-degrees-labor-market/584479/" data-omni-click="inherit">Where Have All the Men Without College Degrees Gone? </a></li>
        
        <li><a href="https://www.theatlantic.com/science/archive/2019/03/indoor-plants-clean-air-best-none-them/584509/" data-omni-click="inherit">A Popular Benefit of Houseplants Is a Myth </a></li>
        
        <li><a href="https://www.theatlantic.com/technology/archive/2019/03/china-has-stopped-accepting-our-trash/584131/" data-omni-click="inherit">Is This the End of Recycling? </a></li>
        
        <li><a href="https://www.theatlantic.com/magazine/archive/2019/04/john-bolton-trump-national-security-adviser/583246/" data-omni-click="inherit">Will John Bolton Bring on Armageddon—Or Stave It Off? </a></li>
        
        <li><a href="https://www.theatlantic.com/ideas/archive/2019/03/surprising-truth-about-roosevelts-new-deal/584209/" data-omni-click="inherit">The New Deal Wasn’t What You Think </a></li>
        
        <li><a href="https://www.theatlantic.com/ideas/archive/2019/03/what-manafort-sentence-reveals/584452/" data-omni-click="inherit">6 Reasons Paul Manafort Got Off So Lightly </a></li>
        
        <li><a href="https://www.theatlantic.com/magazine/archive/2019/04/william-j-burns-putin-russia/583255/" data-omni-click="inherit">How the U.S.-Russian Relationship Went Bad </a></li>
        
    </ol>
</section>

                
                <div class="module module-newsletter-signup">
    <h2 class="module-label">Newsletter Signup</h2>
    <form action="" id="photo-newsletter-signup" data-newsletter-signup="1" method="post" target="_blank">
        <input type="hidden" name="newsletters" value="7581260">
        <input type="hidden" name="newsletterName" value="In Focus">
        <input type="hidden" name="message" value="Newsletter Signup">
        <input type="text" name="email" placeholder="Email address" />
        <input type="submit" value="Go" />

        <label for="sponsored-newsletter">
          <input
            id="sponsored-newsletter"
            type="checkbox"
            name="newsletters"
            
            value="7607877" checked
            
          >
          I want to receive updates from partners and sponsors.
        </label>

        

        <p class="message"></p>
    </form>
</div>

            </div>

        
            <div class="to-top">
                <a href="#article" class="btn" data-omni-click="r'photo',r'below-article',d,r'to-top',r' ',l.pathname">Back to Top</a>
            </div>
            <gpt-ad
                targeting-pos="bottomboxrightphoto"
                class="ad ad-bottomboxrightphoto"
                >
                <gpt-sizeset viewport-size="[0, 0]" sizes="[[300, 250], [320, 50]]"></gpt-sizeset>
            </gpt-ad>
        
    </div>
</div>




        
            
            <div id="janrain-screens" class="profile-screens"></div>
        

        
          <div id="leaflets" class="l-article-recirc__leaflet u-dynamic-content lazyload" data-include="module:theatlantic/js/components/leaflet">
            <span class="u-dynamic-content__loader"></span>
          </div>
        

        
            <footer id="site-footer" class="c-footer lazyload" data-include="module:theatlantic/js/components/footer" role="contentinfo"><div class="c-footer__container"><nav class="c-footer__nav c-footer__nav--top" id="site-footer-top-nav"><ul class="c-footer__list c-footer__list--top"><li class="c-footer__section"><div class="c-footer__section-title">About</div><div class="c-footer__secondary" id="site-footer-about"><ul class="c-footer__list c-footer__list--secondary"><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/history/">Our History</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/masthead/">Staff</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/jobs/">Careers</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/store/">Shop</a></li></ul></div></li><li class="c-footer__section"><div class="c-footer__section-title">Contact</div><div class="c-footer__secondary" id="site-footer-contact"><ul class="c-footer__list c-footer__list--secondary"><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="https://support.theatlantic.com/">Help Center</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/contact/">Contact Us</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/tips/">Send a News Tip</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/advertise/">Advertise</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/press/">Press</a></li></ul></div></li><li class="c-footer__section"><div class="c-footer__section-title">Podcasts</div><div class="c-footer__secondary" id="site-footer-podcasts"><ul class="c-footer__list c-footer__list--secondary"><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/podcasts/radio-atlantic/">Radio Atlantic</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/podcasts/crazygenius/">Crazy/Genius</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/podcasts/the-atlantic-interview/">The Atlantic Interview</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/podcasts/audio-articles/">Audio Articles</a></li></ul></div></li><li class="c-footer__section"><div class="c-footer__section-title">Subscription</div><div class="c-footer__secondary" id="site-footer-subscription"><ul class="c-footer__list c-footer__list--secondary"><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/subscribe/footer-cover/">Purchase</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/subscribe/footer-gift/">Give a Gift</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="https://accounts.theatlantic.com">Manage Subscription</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/app/">Download iOS App</a></li><li class="c-footer__item c-footer__item--top"><a class="c-footer__link" href="/newsletters/sign-up/">Newsletters</a></li></ul></div></li><li class="c-footer__section"><div class="c-footer__section-title">Follow</div><div class="c-footer__secondary" id="site-footer-follow"><ul class="c-footer__list c-footer__list--secondary c-footer__list--follow"><li class="c-footer__item c-footer__item--follow"><a class="c-footer__link" href="https://www.facebook.com/TheAtlantic" title="Like us on Facebook" rel="noopener" target="_blank"><svg class="o-icon c-footer__icon"  ><use xlink:href="#icon-social-facebook"/></svg></a></li><li class="c-footer__item c-footer__item--follow"><a class="c-footer__link" href="https://www.instagram.com/theatlantic" title="Follow us on Instagram" rel="noopener" target="_blank"><svg class="o-icon c-footer__icon"  ><use xlink:href="#icon-social-instagram"/></svg></a></li><li class="c-footer__item c-footer__item--follow"><a class="c-footer__link" href="https://www.youtube.com/user/TheAtlantic" title="Subscribe on YouTube" rel="noopener" target="_blank"><svg class="o-icon c-footer__icon"  ><use xlink:href="#icon-social-youtube"/></svg></a></li></ul><ul class="c-footer__list c-footer__list--secondary c-footer__list--follow"><li class="c-footer__item c-footer__item--follow"><a class="c-footer__link" href="https://twitter.com/TheAtlantic" title="Follow us on Twitter" rel="noopener" target="_blank"><svg class="o-icon c-footer__icon"  ><use xlink:href="#icon-social-twitter"/></svg></a></li><li class="c-footer__item c-footer__item--follow"><a class="c-footer__link" href="https://www.linkedin.com/company/the-atlantic" title="Connect with us on LinkedIn" rel="noopener" target="_blank"><svg class="o-icon c-footer__icon"  ><use xlink:href="#icon-social-linkedin"/></svg></a></li><li class="c-footer__item c-footer__item--follow"><a class="c-footer__link" href="https://flipboard.com/@theatlantic" title="Follow us on Flipboard" rel="noopener" target="_blank"><svg class="o-icon c-footer__icon"  ><use xlink:href="#icon-social-flipboard"/></svg></a></li></ul><ul class="c-footer__list c-footer__list--secondary c-footer__list--follow"><li class="c-footer__item c-footer__item--follow"><a class="c-footer__link" href="/follow-the-atlantic/#follow-rssfeeds" title="Subscribe to our RSS feed" rel="noopener" target="_blank"><svg class="o-icon c-footer__icon"  ><use xlink:href="#icon-social-rss"/></svg></a></li></ul></div></li></ul></nav><div class="c-footer__bottom" id="site-footer-bottom"><div class="c-footer__bottom-info"><nav class="c-footer__nav c-footer__nav--bottom" id="site-footer-bottom-nav"><ul class="c-footer__list"><li class="c-footer__item c-footer__item--bottom"><a class="c-footer__link" href="/privacy-policy/">Privacy&nbsp;Policy</a></li><li class="c-footer__item c-footer__item--bottom"><a class="c-footer__link" href="https://cdn.theatlantic.com/static/front/docs/ads/TheAtlanticAdvertisingGuidelines.pdf?v1">Advertising&nbsp;Guidelines</a></li><li class="c-footer__item c-footer__item--bottom"><a class="c-footer__link" href="/terms-and-conditions/">Terms&nbsp;Conditions</a></li><li class="c-footer__item c-footer__item--bottom"><a class="c-footer__link" href="/responsible-disclosure-policy/">Responsible&nbsp;Disclosure</a></li><li class="c-footer__item c-footer__item--bottom"><a class="c-footer__link js-prefer-hp--us" href="/">U.S.&nbsp;Edition</a></li><li class="c-footer__item c-footer__item--bottom"><a class="c-footer__link js-prefer-hp--world" href="/world/">World&nbsp;Edition</a></li><li class="c-footer__item c-footer__item--bottom"><a class="c-footer__link" href="/site-map/">Site&nbsp;Map</a></li></ul></nav><p class="c-footer__copyright">TheAtlantic.com Copyright (c) 2019 by The Atlantic Monthly Group. All Rights Reserved.</p></div><p class="c-footer__logo"><a class="c-logo" href="/"><span class="u-element-invisible c-logo__name">The Atlantic</span> <svg class="o-icon"  viewBox="0 0 256 84"  ><use xlink:href="#icon-logo-atl"/></svg></a></p></div></div></footer>
        

    
</div>




<script type="gpt-config">
    {"ix": true, "lazy_load": 3, "datePublished": "2019-03-08T14:26:13-05:00", "amazon": "3239", "globals": {"src": "article", "watson_categories": ["/pets/dogs", "/sports/snowboarding"], "title": "photos-of-the-week-frozen-road-broadway-gorilla-penguin-swing-set", "report": null, "type": "photo", "grapeshot_segments": ["an_bankofamerica_negative", "an_facebook_safe", "an_fidelity_safe", "an_jpmc_safe", "an_ms_safe", "gs_home_pets", "gs_sport", "gs_home", "gs_politics_misc", "gs_event_blackfriday", "gs_politics"], "id": 584431}, "prebid": "https://www.theatlantic.com/packages/adsjs/prebid.js", "perfUrl": "https://data-cdn.theatlantic.com/perf.gif", "zone": "/4624/TheAtlanticOnline/channel_international", "adtest_domain": "dctestsite", "patchEventHandlers": true, "criteo": true, "outofpage": true, "gdpr": false}
</script>

<div id="outofpage"></div>




    
    <script>
      (function() {
        function hasOwnProp(o, k) {
          return Object.prototype.hasOwnProperty.call(o, k);
        }

        function entries(o, cb) {
          for (var k in o) { if (hasOwnProp(o, k)) cb(k, o[k]); }
        }

        var staticHashMap = {"components/homepage": {"ww3": "ebaf3c2245ed", "cover-story": "e2b23e555d09", "offlede": "2086c5f61e5c", "feature": "ef245ac7ce39", "recent-issues": "d8e6a10bd85b", "podcasts": "ff8b3ce57a74", "writers": "ca3f7b4c996b", "latest-videos": "bef584255186", "popular": "0523c3cbb9e8", "subscription": "2bb60bc3bce6", "sections": "348f11364511", "latest": "618fe7a8cdc0"}, "components/ads": {"box": "496a5e78cfbc", "standard-wide": "c483343afe68", "article-mobile": "c8489c6cbd86", "native-mobile": "18962ea81105", "article-house-mobile": "5d7718a2bd0c", "bottompersistent": "ac0509996c0b", "rail": "7abd43a0a096", "native-desktop": "86357b89e4fb", "article-house-desktop": "7449bb57673c", "cineflex": "ba177f6957fa", "logo": "b97423a7913e", "sponsored-links": "aded2f3d99d7", "pixel": "00dbda0d88b0", "article-desktop": "866cf0506473", "native": "5a853ec746e7"}, "components/ideas": {"article-meta": "d3919defa1a7", "dateline": "23bc8c8f8c0d", "dek": "b3ef4307dc99", "article-author": "605acf3da32e", "article-header": "cba6b260e655"}, "sw": {"sw": "8530eaa68450"}, "juvenile-justice": {"juvenile-justice.min": "6f3ca222a9a6"}, "js/bundle/components": {"recirc-content": "6861870e1244", "footer": "3e4b7901d272", "article-video": "14f12f496b1f", "share-social": "920816989b56", "leaflet": "08511aa9bcc5", "most-popular": "4cf796960f5e", "navigation": "7e0d3c75a3f6", "related-video": "89d1a02ed700"}, "js/bundle/pages": {"year-in-review": "3241f872c276"}, "components/year-in-review-2018": {"king": "b6565fd6204a", "initiatives": "60b6d197a9e9", "democracy": "3fa6e8516768", "header": "71e6ef4db904", "people": "c60e4f77f892", "studios": "2a52992070c2", "podcasts": "0b59341faf6b", "cta-buttons": "88ea2f970728", "more-stories": "bc42c99903dd", "nav": "36146012cbe3", "stories": "0282e0226766", "intro": "1e34f7ddffa0", "events": "0df290b48508"}, "components": {"breaking-news": "ec5a2ae3611d", "ad": "ac2b0f00cfdd", "share-social": "ab1e5c655020", "audio-embed": "99047955e3cc", "rubric": "e48916c2438a", "sentinel": "05714e19cc77", "newsletter": "472e68c5fc33", "foundation-line": "000f34f96e76", "letters-cta": "aa4eb1318560", "story": "d117e8f0d35b", "byline": "9dcf2ce5b50f", "icons": "0c5c28b4f49a", "dateline": "523100f5afdb", "nav-logo": "e771dd1f249b", "paywall": "fb45377e2bb7", "book-review": "fdd52d5b4603", "card": "b767742bb64f", "lead-image": "cb418ce25e55", "input": "2b20170fa7dc", "tease-article": "e846fdab8980", "audio-emblazon": "8fd80cdd89f4", "article-header": "7d35e898ed25", "books-in-review": "18d74f71b0c0", "picture": "18027a2f7991", "lead-cinemagraph": "d6d5def44887", "hero": "12756aebf7ad", "bottom-persistent-ad": "22f468212dee", "story-strip": "857a05449c27", "nav-menu-list": "7ceb8703cc25", "social-links": "4a1f3d428dbd", "review": "e1293b006c89", "skip-link": "28f56f682930", "lead-embed": "bc35fca0cbd3", "categories": "85422111c339", "icon": "6f545c12d762", "article-writer": "b776cea25ae1", "recirc-item": "96c1af2b13ae", "trump50-hp": "61c25193954a", "section-header": "309909e42a78", "article-meta": "dd8570491b6f", "recirc-content": "156270ba2efa", "footer": "03163665b99c", "article-video": "482a04d20d03", "button": "32046ce97854", "tease-audio": "aa7e152fa37d", "_preview": "f714f3dd5e15", "nav-desktop-search": "897d86cf7448", "leaflet": "6481b4aa30c4", "most-popular": "83c899565eb0", "dek": "caa04aef968e", "editors-note": "777c69ce10af", "river": "eece6c186321", "navigation": "b3f3bfb3f2c5", "nav-mobile-search": "abd40bd26354", "related-video": "784a15a4cb85"}, "js/bundle": {"homepage": "898468c70e3b"}, "js": {"hippo": "91a9de98dca5", "legacy": "96a6706e50cd", "system": "85f5cf2c11e5", "bundle": "e006ebadda77", "polyfill": "3064abc2580b"}, "js/bundle/utils": {"debug": "6b1a007d1eaa", "iframe-resizer": "209793ea174d"}};
        var systemjsMap = {};

        entries(staticHashMap, function(dir, files) {
          entries(files, function(basename, hash) {
            var path = 'dist/theatlantic/' + dir + '/' + basename;
            systemjsMap[path + '.js'] = path + '.' + hash + '.js';
          });
        });

        SystemJS.config({
          baseURL: "https://cdn.theatlantic.com/assets/static/a/frontend/",
          map: systemjsMap
        });
        SystemJS.set('pageInfo', SystemJS.newModule(window.Atlantic.page_info));
        SystemJS.set('profilesConfig', SystemJS.newModule({
          profiles_url: window.Atlantic.PROFILES_URL,
          janrain: window.Atlantic.SESSION_URLS
        }));
      }());
    </script>

    <script type="text/javascript" src="https://cdn.theatlantic.com/assets/static/a/theatlantic/js/footer.min.9563b63c62bb.js" charset="utf-8"></script>



<script type="text/javascript" src="https://cdn.theatlantic.com/assets/static/a/theatlantic/js/photo-article.2bcaccf82c47.js" charset="utf-8"></script>




<!-- Scroll -->
<script>
  (function(a,c,d,e){if(!a[c]){var b=a[c]={};b[d]=[];b[e]=function(a){b[d].push(a)}}})(window,'Scroll','_q','do');
  Scroll.config = {
    orientation: 'bottom',
    detected: document.cookie.indexOf("scroll0=") > -1
  };
</script>
<script async src="https://static.scroll.com/js/scroll.js"></script>







<script type="text/javascript">
    /**
     * Route blockers
     */
     (function(){
        function redirect() {

            // Chrome 54-57 has a bug with its extensions that breaks
            // whitelisting on the first pageview when sent through
            // a shortened url. We can't detect shorturls reliable,
            // so let it slide.
            var chromeFreePV = (function(){
                var host = window.location.protocol + "//" + window.location.host;
                var isFirstPV = (document.referrer.indexOf(host) !== 0);
                var isChrome = (window.navigator.userAgent.match(/Chrome\//) !== null);
                return isFirstPV && isChrome;
            })();

            var isCrawler = (function(){
                var botsRe = "(Googlebot|Bingbot|Yahoo! Slurp|DuckDuckBot|facebookexternalhit)";
                return (window.navigator.userAgent.match(botsRe) !== null);
            }());

            if ((navigator.onLine === false) ||
                isCrawler ||
                chromeFreePV ||
                Atlantic.User.isMember() ||
                Atlantic.User.isAdFree({strict: false})) {
                    return;
            }
            var next = encodeURIComponent(window.location.href);
            var url = window.location.protocol + "//" + window.location.host + "/please-support-us/?next=";
            window.location.href = url + next;
        };

        // Prevent race condition
        if (Modernizr["blocker-enabled"]) {
            redirect();
        } else {
            $(window).one("blocker-detected", redirect);
        }
     })();
</script>







<!-- 2019-03-10T04:23:34.782974 -->
<!-- PAGE_COMPLETED -->

</body>
</html>
`

func TestInFocusFeedBuilder(t *testing.T) {
	node, _ := goquery.NewDocumentFromReader(strings.NewReader(infocusFile))
	handler := InFocusFeed{}

	_, err := handler.ParseFeed("", "id", 12345, node)
	if err != nil {
		t.Error("Should parse the file")
	}
}
