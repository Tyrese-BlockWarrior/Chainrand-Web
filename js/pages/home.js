(function () {
	
	var vertCode = [
		'#ifdef GL_ES', 
		'precision mediump float;', 
		'#endif', 
		'attribute vec2 aPos;', 
		'void main() {', 
			'gl_Position = vec4(aPos, 0.0, 1.0);', 
		'}'].join('\n');

	var fragCode = [
		'// Author @patriciogv - 2015', 
		'// Title: DeFrag', 
		'// Adapted from https://thebookofshaders.com/edit.php#10/ikeda-04.frag', 
		'#ifdef GL_ES', 
		'precision mediump float;', 
		'#endif', 
		'uniform float u[5];', 
		'float random (float x) { return fract(sin(x)*1e4); }', 
		'float random (vec2 _st) { return fract(sin(dot(_st.xy, vec2(12.9898,78.233)))* 43758.5453123);}', 
		'void main() {', 
			'vec2 u_resolution = vec2(u[0], u[1]);', 
			'vec2 u_mouse = vec2(u[2], u[3]);', 
			'float u_time = u[4];', 
			'vec2 st = gl_FragCoord.xy/u_resolution.xy;', 
			'st.x *= u_resolution.x/u_resolution.y;', 
			'vec2 grid = vec2(100.0,100.0);', 
			'st *= grid;', 
			'vec2 ipos = floor(st);', 
			'vec2 vel = vec2(u_time*10.);', 
			'vel *= vec2(-1.,0.);', 
			'vel *= (step(1., mod(ipos.y,2.0))-0.5)*2.; // Opposite directions', 
			'vel *= random(ipos.y); // random speed', 
			'vec3 color = vec3(1.0);', 
			'color *= step(grid.y,ipos.y);', 
			'color += step(0.0,ipos.y);', 
			'color = clamp(color,vec3(0.),vec3(1.));', 
			'float fadeFactor = abs(u_mouse.y/u_resolution.y-0.5);', 
			'fadeFactor = fadeFactor + abs(u_mouse.x/u_resolution.x-0.5);', 
			'fadeFactor = 0.5 + clamp(fadeFactor * 0.5, 0.1, 0.5) * 0.5;', 
			'color *= random(floor(st+vel));', 
			'color = smoothstep(0.0, fadeFactor, color*color);', 
			'color = step(fadeFactor,color); // threshold',
			'// Margin', 
			'float margin = 0.2;', 
			'color *= step(margin,fract(st.x+vel.x))*step(margin,fract(st.y+vel.y));', 
			'gl_FragColor = vec4(1.90-color,1.0);', 
		'}'].join('\n');

	var $page, $window, $canvas;

	var canvas, gl, uUniform, uBuffer;
	var canvasWidth, canvasHeight, canvasMouseX, canvasMouseY;
	var prevVW, prevVH;
	var animFrame, animTicker = 0, animTickerMod = 4;
	var posBuffer, uvBuffer, posVBO, uvVBO;

	var visible = 0;

	var contentCtx = ['div', 
		['div', {class: 'splash'}, 
			['canvas', {class: 'bg'}],
			['div', {class: 'front'}, 
				['div', {class: 'title'}, 'Fair Randomness Anywhere'],
				['div', {class: 'tagline'}, 
					'Generate verfiable random numbers efficiently without limits'],
				['div', {class: 'buttons'},
					['a', {class: 'mint', href: '#/mint'}, 
						['div','Mint'], ['div', {class: 'b-0'}], ['div', {class: 'b-1'}]
					],
					['div', {class: 'spacer'}],
					['a', {class: 'github', href: githubURL, target: '_blank'}, 
						['div','Github'], ['div', {class: 'b-0'}], ['div', {class: 'b-1'}]
					],
				]
			]
		]
	]; 

	var t0 = 'Fairness Without Compromises';
	var t1 = ['div', 
		['p', 'Random number generators are prevalent in our everyday lives. '],
		['p', 'It is important for RNGs to be unbiased and incorruptible.'],
		['p',
			'However, RNG seeds can be cherry-picked to achieve ',
			'unfair outcomes in subtle ways. ',
			'This is a common problem in many fields ',
			'known as Statistical P-Hacking, ', 
			'which is very hard to detect, or disprove.'
		],
		['p',
			'With Chainrand, it is now easy to use ', 
			'Chainlink Verifiable Random Functions ',
			'for flexible, efficient computations in off-chain code ',
			'in a verfiably fair way. '
		]
	];
	var introCtx = ['div', {class: 's-intro'}, 
		['div', {class: 'desktop'},
			['table', ['tr',
				['td', {class: 't-0'},
					['div', {class: 'title'}, t0.replace(/\s/g, '<br>')], 
				],
				['td', {class: 't-1'},
					['div', {class: 'text'}, t1]
				]
			]]
		], 
		['div', {class: 'mobile'}, 
			['div', {class: 'title'}, t0],
			['div', {class: 'text'}, t1] 
		]
	]; 
	contentCtx.push(introCtx);

	var t0 = ['div', {class: 'field'},
		['img', {src: 'assets/home-academia.png'}],
		['div', {class: 'text'}, 'Research']	
	];
	var t1 = ['div', {class: 'field'},
		['img', {src: 'assets/home-finance.png'}],
		['div', {class: 'text'}, 'Finance']	
	];
	var t2 = ['div', {class: 'field'},
		['img', {src: 'assets/home-gaming.png'}],
		['div', {class: 'text'}, 'Gaming']	
	];
	var t3 = ['div', {class: 'field'},
		['img', {src: 'assets/home-art.png'}],
		['div', {class: 'text'}, 'Art']	
	];
	var ts = ['div', {class: 'spacer'}]

	var applicationsCtx = ['div', {class: 's-applications'},
		['div', {class: 'title'}, 'Applications'],
		['div', {class: 'desktop'},
			t0, ts, t1, ts, t2, ts, t3
		],
		['div', {class: 'mobile'},
			['div', t0, ts, t1],
			['div', t2, ts, t3]
		]
	];
	contentCtx.push(applicationsCtx);

	var t0 = ['div', 
		['div', {class: 'title'}, 'Getting Started'], 
		['div', {class: 'info'}, 
			'Chainrand provides determinstic cryptographically secure RNG SDKs ',
			'for several languages. More languages will be added over time.'
		],
		['div', {class: 'sdks'},
			['a', {class: 'lang', href: githubURL + '/chainrand-js', target: '_blank'}, 
				['img', {src: 'assets/home-lang-js.svg'}],
				['div', {class: 'text'}, 'Javascript']
			],
			['a', {class: 'lang', href: githubURL + '/chainrand-py', target: '_blank'}, 
				['img', {src: 'assets/home-lang-py.svg'}],
				['div', {class: 'text'}, 'Python']
			],
			['a', {class: 'lang', href: githubURL + '/chainrand-cpp', target: '_blank'}, 
				['img', {src: 'assets/home-lang-cpp.svg'}],
				['div', {class: 'text'}, 'C++']
			],
		],
		['div', {class: 'info'}, 
			'For a more comprehensive process, checkout our demo generative NFT project. '
		],
		['div', {class: 'demo'}, 
			['a', {href: githubURL + '/chainrand-demo', target: '_blank'},
				'Link to Demo project ', ['i', {class: 'icon-link-ext'}]
			]
		],
	];
	var arrowCtx = ['div', {class: 'arrow'}, ['i', {class:'icon-angle-down'}]];
	var t1 = ['div', {class: 'process'}, 
		['div', {class: 'step'},
			'Write off-chain determinstic RNG code with a Chainrand SDK, ', 
			'and publish it.'
		], 
		arrowCtx,
		['div', {class: 'step'},
			'Mint a Chainrand NFT to permanently tie the code to an ',
			'unpredictable, immutable RNG seed generated with Chainlink VRF.'
		],
		arrowCtx,
		['div', {class: 'step'},
			'Run the code, publish the results, and update the Chainrand NFT.'
		]
	];
	
	var startedCtx = ['div', {class: 's-started'},
		['div', {class: 'desktop'}, 
			['table', ['tr',
				['td', {class: 't-0'}, t0], ['td', {class:'t-1'}, t1],
			]]
		],
		['div', {class: 'mobile'}, 
			t0, 
			['div', {class: 'subtitle'}, 'The Process'], 
			t1
		]
	];
	contentCtx.push(startedCtx);
	// 	['div', {class: 'applications'}, // white bg
	// 		['div', {class: 'title'}, 'Applications'],
			
	// 	], 
	// 	['div', {class: 'started'}, 
	// 		['div', {class: 'title'}, 'Getting started'],

	// 	]
	// ];

	// Write off-chain determinstic RNG code with a Chainrand SDK.  
	// Generate RNG seed for the RNG code with Chainrand.
	// Reveal the results of the RNG code with Chainrand.

	var createVBO = function (program, data, dims, name) {
		var vbo = gl.createBuffer();
		gl.bindBuffer(gl.ARRAY_BUFFER, vbo);
		gl.bufferData(gl.ARRAY_BUFFER, data, gl.STATIC_DRAW);
		var attr = gl.getAttribLocation(program, name);
		gl.vertexAttribPointer(attr, dims, gl.FLOAT, false, 0, 0);
		gl.enableVertexAttribArray(attr);
		gl.bindBuffer(gl.ARRAY_BUFFER, null);
		return vbo;
	};

	var canvasTime = 0.0;

	var requestAnimFrame = (function(){
		return window.requestAnimationFrame ||
			window.webkitRequestAnimationFrame ||
			window.mozRequestAnimationFrame ||
			function (callback) {
				return window.setTimeout(callback, 1000 / 60);
			};
	})();
	var cancelAnimFrame = (function(){
		return window.cancelAnimationFrame ||
			window.webkitCancelAnimationFrame ||
			window.mozCancelAnimationFrame ||
			function (a) {
				return window.clearTimeout(a);
			};
	})();
	
	var stepCanvas = function () {
		
		if (animTicker < 1) {
			var vw = $window.width();
			var vh = $window.height();
			if (vw != prevVW || vh != prevVH) {
				canvas.height = canvasHeight = Math.max(350, vh * 0.7)
				canvas.width = canvasWidth = vw;
				gl.viewport(0, 0, canvasWidth, canvasHeight);
				prevVH = vh; 
				prevVW = vw;
			}
			uBuffer[0] = canvasWidth; 
			uBuffer[1] = canvasHeight; 
			uBuffer[2] = canvasMouseX; 
			uBuffer[3] = canvasMouseY; 
			uBuffer[4] = canvasTime;
			canvasTime += 0.005 * animTickerMod; 
			gl.uniform1fv(uUniform, uBuffer);
			// gl.clearColor(0.0, 0.0, 0.0, 1.0);
			gl.clear(gl.COLOR_BUFFER_BIT);
			gl.drawArrays(gl.TRIANGLES, 0, 6);
		}
		animTicker = (animTicker + 1) % animTickerMod;

		if (visible)
			animFrame = requestAnimFrame(stepCanvas);
	}

	var setupCanvas = function () {
		$canvas = $page.find('canvas'); 
		canvas = $canvas[0];
		gl = canvas.getContext('webgl');
		
		var p = new Float32Array([-1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0]);
		
		posBuffer = p;
		
		var vertShader = gl.createShader(gl.VERTEX_SHADER);   
		gl.shaderSource(vertShader, vertCode);   
		gl.compileShader(vertShader);   

		var fragShader = gl.createShader(gl.FRAGMENT_SHADER);   
		gl.shaderSource(fragShader, fragCode);   
		gl.compileShader(fragShader);   

		var shaderProgram = gl.createProgram();   
		gl.attachShader(shaderProgram, vertShader);   
		gl.attachShader(shaderProgram, fragShader);   
		gl.linkProgram(shaderProgram);   
		gl.useProgram(shaderProgram);   
		
		posVBO = createVBO(shaderProgram, posBuffer, 2, 'aPos');

		uUniform = gl.getUniformLocation(shaderProgram, 'u');   

		uBuffer = new Float32Array(5);

		$canvas.on('mousemove', function (e) {
			canvasMouseX = e.pageX;
			canvasMouseY = e.pageY;
		})
	}

	var init = function (sel) {
		$page = $(sel).html(HTML(contentCtx));
		$window = $(window);
		setupCanvas();
	}

	var sync = function (sel) {

	}

	var show = function () {
		if (!visible) {
			visible = 1;
			animFrame = requestAnimFrame(stepCanvas);	
		}
	}

	var hide = function () {
		if (visible) {
			cancelAnimFrame(animFrame);
			visible = 0;
		} 
	}

	Layout.add({
		label: 'home', 
		init: init, 
		sync: sync,
		show: show, 
		hide: hide
	});
})()