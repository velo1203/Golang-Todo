import React, { useEffect, useRef } from "react";
import { StyledPopupWrapper } from "../../../style/common/StyledPopupWrapper";

function PopupWrapper({ children, onOutsideClick = () => {} }) {
    const wrapperRef = useRef(null);

    useEffect(() => {
        function handleWrapperClick(event) {
            // 클릭된 요소가 wrapperRef 자체일 때 onOutsideClick 호출
            if (wrapperRef.current && wrapperRef.current === event.target) {
                if (typeof onOutsideClick === 'function') {
                    onOutsideClick();
                }
            }
        }

        // 이벤트 리스너를 wrapperRef에 추가
        if (wrapperRef.current) {
            wrapperRef.current.addEventListener("click", handleWrapperClick);
        }

        return () => {
            // 이벤트 리스너 정리
            if (wrapperRef.current) {
                wrapperRef.current.removeEventListener("click", handleWrapperClick);
            }
        };
    }, [onOutsideClick]);

    return <StyledPopupWrapper ref={wrapperRef}>
        {children}
    </StyledPopupWrapper>;
}

export default PopupWrapper;
