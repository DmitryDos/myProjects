let posX = 0;
let posY = 0;
let posXLast = 0;
let posYLast = 0;
let body = document.querySelector("body");
let colors = ["red", "orange", "lightgreen", "yellow"];
let colorNumb = 0;
let type = 0;
body.addEventListener("mousemove", function(e){
	posX = e.clientX;
	posY = e.clientY;
	
});
body.addEventListener("click", function(){
		type = 1 - type;
})
makeBullets();
function makeBullets(){
	let pos1 = posX;
	let pos2 = posY;
	let size = Math.sqrt(Math.abs(pos1 - posXLast) * Math.abs(pos1 - posXLast) + Math.abs(pos2 - posYLast) * Math.abs(pos2 - posYLast) + 400) / 4;
	posXLast = pos1;
	posYLast = pos2;
	makeBullet(0, size, pos1, pos2, colors[colorNumb]);
	makeBullet(1, size, pos1, pos2, colors[colorNumb]);
	makeBullet(2, size, pos1, pos2, colors[colorNumb]);
	makeBullet(3, size, pos1, pos2, colors[colorNumb]);
	colorNumb = (colorNumb + 1) % 4;
	setTimeout(makeBullets, 50);
}
function makeBullet(dir, size, pos1, pos2, color){
	let el = document.createElement("div");
	el.style.position = "absolute";
	el.style.transform = "translate(-50%, -50%)";
	el.style.left = pos1;
	el.style.top = pos2;
	el.style.height = size + "px";
	el.style.width = size + "px";
	if(type == 0){
		el.style.borderRadius = "50%";
	}
	el.style.background = color;
	el.style.pointerEvents = "none";
	body.appendChild(el);
	setTimeout(moveBullet, 0, el, dir, pos1, pos2, 1000);
}
function moveBullet(el, dir, pos1, pos2, t){
	if(t > 0){
		t -= 10;
		if(dir == 0){
			pos2 += 1;
		}
		else if(dir == 1){
			pos1 += 1;
		}
		else if(dir == 2){
			pos2 -= 1;
		}
		else if(dir == 3){
			pos1 -= 1;
		}
		el.style.top = pos2;
		el.style.left = pos1;
		el.style.opacity = (t / 1000);
		setTimeout(moveBullet, 10, el, dir, pos1, pos2, t);
	}
	else{
		el.remove();
	}
}
