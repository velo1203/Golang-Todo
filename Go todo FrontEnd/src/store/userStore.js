import create from 'zustand';

const useAuthStore = create((set) => ({
  userToken: localStorage.getItem('userToken') || null, // 로컬 스토리지에서 토큰 가져오기
  isLoggedIn: !!localStorage.getItem('userToken'), // 로그인 상태 초기화
  setUserToken: (token) => {
    localStorage.setItem('userToken', token); // 여기에서 로컬 스토리지에 토큰 저장
    set({ userToken: token ,isLoggedIn: true});
  },
  logout: () => {
    localStorage.removeItem('userToken'); // 로컬 스토리지에서 토큰 제거
    set({ userToken: null, isLoggedIn: false }); // 로그아웃 상태 설정
  }
}));

export default useAuthStore;
