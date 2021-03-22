netlify:
	find api/netlify/ -name "*.go" -exec ./netlify-functions.sh {} \;; \
	sed -i 's/API_URL/.netlify\/functions/g' ./public/index.html
vercel:
	sed -i 's/API_URL/api\/vercel/g' ./public/index.html