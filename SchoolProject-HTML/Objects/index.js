let VW = window.innerWidth;
let VH = window.innerHeight;
let resizeFunction = function(){
	VH = window.innerHeight;
	VW = window.innerWidth;
}
window.onresize = resizeFunction;
onscroll = function(){
	let el;
	
	el = document.querySelector(".it1");
	if(el.getBoundingClientRect().top + el.getBoundingClientRect().height < VH * 1){
		el.className = "introduceText it1 it12";
	}
	else{
		el.className = "introduceText it1";
	}
	
	el = document.querySelector(".it2");
	if(el.getBoundingClientRect().top + el.getBoundingClientRect().height < VH * 1){
		el.className = "introduceText it2 it22";
	}
	else{
		el.className = "introduceText it2";
	}
	
	el = document.querySelector(".it3");
	if(el.getBoundingClientRect().top + el.getBoundingClientRect().height < VH * 1){
		el.className = "introduceText it3 it32";
	}
	else{
		el.className = "introduceText it3";
	}
	
	el = document.querySelector(".logicTitle");
	if(el.getBoundingClientRect().top < VH){
		document.querySelector(".introducePage").style.background = "rgb(255, 255, 255)";
		document.querySelector(".logicPage").style.background = "rgb(255, 255, 255)";
	}
	else{
		document.querySelector(".introducePage").style.background = "rgb(0, 0, 0)";
		document.querySelector(".logicPage").style.background = "rgb(0, 0, 0)";
	}
	
	if(el.getBoundingClientRect().top + el.getBoundingClientRect().height < 0){
		document.querySelector(".logicPictureContainer").style.opacity = "1";
	}
	else{
		document.querySelector(".logicPictureContainer").style.opacity = "0";
	}
	
	for(let i = 1; i < 6; i++){
		el = document.querySelector(".ltb" + i);
		if(el.getBoundingClientRect().top + el.getBoundingClientRect().height < VH && el.getBoundingClientRect().top + el.getBoundingClientRect().height > 0){
			let pos = el.getBoundingClientRect().top + el.getBoundingClientRect().height;
			let d = 2 * (VH * 0.5 - pos) / VH;
			document.querySelector(".lt" + i).style.transform = "translate("+(150 * d * d * d)+"%, 0)";
		}
		else{
			document.querySelector(".lt" + i).style.transform = "translate(-150%, 0)";
		}
	}
	el = document.querySelector(".logicPage");
	if(-VH < el.getBoundingClientRect().top + el.getBoundingClientRect().height && el.getBoundingClientRect().top + el.getBoundingClientRect().height < 0){
		document.querySelector(".logicPictureContainer").style.marginLeft = 100 * (0 - el.getBoundingClientRect().top - el.getBoundingClientRect().height) / VH + "vw";
	}
	else if(-VH > el.getBoundingClientRect().top + el.getBoundingClientRect().height){
		document.querySelector(".logicPictureContainer").style.marginLeft = "100vw";
	}
	else{
		document.querySelector(".logicPictureContainer").style.marginLeft = "0vw";
	}
	
};
window.addEventListener("scroll", onscroll);

let linksButton = document.querySelector(".menuLinksContainer");
let linksNav = document.querySelector(".linksNav");
let isLinksOpened = false;
linksButton.addEventListener("click", function(){
	if(isLinksOpened == false){
		isLinksOpened = true;
		linksNav.style.top = "5px";
	}
	else{
		isLinksOpened = false;
		linksNav.style.top = "-100px";
	}
});

let header = document.querySelector("header");
header.addEventListener("mouseleave", function(){
	isLinksOpened = false;
	linksNav.style.top = "-20vh";
});
let isBgColoring = false;
let BGContainer = document.querySelector(".BGContainer");
let BGColoring = function(){
	BGContainer.removeEventListener("click", BGColoring);
	document.querySelector(".helpText1").style.opacity = 0;
	document.querySelector(".colorBGMid").style.opacity = 1;
	setTimeout(function(){
		document.querySelector(".colorBG").style.transform = "translate(-50%, -50%) scale(2) rotate(45deg)";
		document.querySelector(".helpText1").style.display = "none";
	}, 500);
}
BGContainer.addEventListener("click", BGColoring);

document.querySelector(".mi1").addEventListener("click", function(){
	document.querySelector(".introduceTitle").scrollIntoView({behavior: "smooth", block: "center"});
});
document.querySelector(".mi2").addEventListener("click", function(){
	document.querySelector(".logicTitle").scrollIntoView({behavior: "smooth", block: "center"});
});
document.querySelector(".mi3").addEventListener("click", function(){
	document.querySelector(".phisicTitle").scrollIntoView({behavior: "smooth", block: "center"});
});
document.querySelector(".mi4").addEventListener("click", function(){
	document.querySelector(".statsTitle").scrollIntoView({behavior: "smooth", block: "center"});
});
document.querySelector(".returnButton").addEventListener("click", function(){
	window.scrollTo(0, 0);
});
function removeExtraPicture(){
	document.querySelector(".extraPictureContainer").style.display = "none";
	document.querySelector("body").style.overflow = "scroll";
	window.removeEventListener("click", removeExtraPicture);
}
for(let i = 1; i <= 4; i++){
	document.querySelector(".pi" + i).addEventListener("click", function(){
		document.querySelector("body").style.overflow = "hidden";
		document.querySelector(".extraPicture").src = "Objects/pi"+i+".jpg";
		document.querySelector(".extraPictureContainer").style.display = "block";
		window.addEventListener("click", removeExtraPicture, true);
	});
}

document.querySelector(".sb0").addEventListener("click", function(){
	document.querySelector(".prizeCounter").innerHTML = "1/5";
	document.querySelector(".sb0").innerHTML = "<p>1&crarr;5</p>";
	document.querySelector(".sb1").style.display = "block";
});
document.querySelector(".sb1").addEventListener("click", function(){
	document.querySelector(".prizeCounter").innerHTML = "2/5";
	document.querySelector(".sb1").innerHTML = "<p>2&#8965;5</p>";
	document.querySelector(".sb2").style.display = "block";
});
document.querySelector(".sb2").addEventListener("click", function(){
	document.querySelector(".prizeCounter").innerHTML = "3/5";
	document.querySelector(".sb2").innerHTML = "<p>3&#8853;5</p>";
	document.querySelector(".sb3").style.display = "block";
});
document.querySelector(".sb3").addEventListener("click", function(){
	document.querySelector(".prizeCounter").innerHTML = "4/5";
	document.querySelector(".sb3").innerHTML = "<p>4&#8788;5</p>";
	document.querySelector(".sb4").style.display = "block";
});
document.querySelector(".sb4").addEventListener("click", function(){
	document.querySelector(".prizeCounter").style.display = "none";
	document.querySelector(".prizeButton").style.display = "block";
	document.querySelector(".sb4").innerHTML = "<p>5/5</p>";
	document.querySelector(".sb4").style.pointerEvents = "none";
	document.querySelector(".sb4").style.opacity = "0";
	document.querySelector(".sb0").remove();
	document.querySelector(".sb1").remove();
	document.querySelector(".sb2").remove();
	document.querySelector(".sb3").remove();
});
