netlify:
	find .netlify/ -name "*.go" -exec ./netlify-functions.sh {} \;; \
	sed -i 's/API_URL/.netlify\/functions/g' ./public/index.html
vercel:
	cp -r functions/vercel/ api/; \
	sed -i 's/API_URL/api/g' ./public/index.html