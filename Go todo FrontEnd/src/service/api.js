import axios from 'axios';
import useAuthStore from '../store/userStore';

// Axios 
const apiClient = axios.create({
    headers: {
      'Content-Type': 'application/json'
    }
  });
  
  // 요청 인터셉터 추가
  apiClient.interceptors.request.use(config => {
    // Zustand 스토어에서 토큰 가져오기
    const token = useAuthStore.getState().userToken;
  
    // 토큰이 있다면 헤더에 추가
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  
    return config;
  }, error => {
    // 요청 에러 처리
    return Promise.reject(error);
  });
  
  export default apiClient;