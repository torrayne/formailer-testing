netlify:
	find api/netlify -name "*.go" -exec ./build-netlify-function.sh {} \;; \
	sed -i 's/API_URL/.netlify\/functions/g' ./public/index.html
vercel:
	cp -r functions/vercel/ api/vercel; \
	sed -i 's/API_URL/api/g' ./public/index.html