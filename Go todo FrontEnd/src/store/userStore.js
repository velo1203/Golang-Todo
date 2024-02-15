import create from 'zustand';
import { persist } from 'zustand/middleware';

const useAuthStore = create(persist(
  (set) => ({
    userToken: null, // 초기 토큰 상태
    isLoggedIn: false, // 초기 로그인 상태
    setUserToken: (token) => {
      set({ userToken: token, isLoggedIn: !!token }); // 토큰 설정 및 로그인 상태 업데이트
    },
    logout: () => {
      set({ userToken: null, isLoggedIn: false }); // 로그아웃 상태 설정
    },
  }),
  {
    name: 'auth', // localStorage에 저장될 때 사용될 키 이름
    getStorage: () => localStorage, // 사용할 스토리지 객체 지정
  }
));

export default useAuthStore;
