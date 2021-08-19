```css


{
	width:100px;
	height:100px;
	border-radius:50%;
	background-image: radial-gradient(#0404B4, #08088A, #0B0B3B, #0B0B61);
	animation: myfirst 5s linear 0s infinite;
}

@keyframes myfirst
{
	0%   {background-image: radial-gradient(#0404B4, #08088A, #0B0B3B, #0B0B61);transform:rotate(0deg);width:100px;height:100px;}
	25%  {background-image: radial-gradient(#0101DF, #0404B4, #08088A, #0B0B3B);transform:rotate(90deg);width:125x;height:125x;}
	50%  {background-image: radial-gradient(#0000FF, #0101DF, #0404B4, #08088A);transform:rotate(180deg);width:175px;height:175px;}
	100% {background-image: radial-gradient(#2E2EFE, #0000FF, #0101DF, #0404B4);transform:rotate(360deg);width:200px;height:200px;}
}
```