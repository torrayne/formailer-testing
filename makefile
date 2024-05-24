netlify:
	find api/netlify -name "*.go" -exec ./build-netlify-function.sh {} \;
vercel:
	cp -r functions/vercel/ api/vercel