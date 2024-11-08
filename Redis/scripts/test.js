import http from 'k6/http'

export let options = {
    vus: 10,
    duration: '1s'
}

export default function() {
    http.get('http://host.docker.internal:8000/GetProduct')
}