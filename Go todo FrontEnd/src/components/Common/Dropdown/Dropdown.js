import React, { useState, useEffect,useRef} from 'react';
import { StyledIcon } from '../../../style/common/StyledIcon';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEllipsisVertical } from '@fortawesome/free-solid-svg-icons';
import { StyledDropdown, StyledDropdownItem, StyledDropdownMenu } from '../../../style/common/StyledDropdown';

function Dropdown() {
    const [isOpen, setIsOpen] = useState(false);
    const dropdownRef = useRef(null);

    const toggleDropdown = () => setIsOpen(!isOpen);

    const handleClickOutside = (event) => {
        if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
            setIsOpen(false);
        }
    };

    useEffect(() => {
        document.addEventListener('mousedown', handleClickOutside);
        return () => {
            document.removeEventListener('mousedown', handleClickOutside);
        };
    }, []);

    return (
        <StyledDropdown ref={dropdownRef}>
            <StyledIcon onClick={toggleDropdown}>
            <FontAwesomeIcon icon={faEllipsisVertical}/>
            </StyledIcon>

            {isOpen && (
                <StyledDropdownMenu>
                    <StyledDropdownItem >Delete</StyledDropdownItem>
                    <StyledDropdownItem >Edit</StyledDropdownItem>
                </StyledDropdownMenu>
            )}
        </StyledDropdown>
    );
}

export default Dropdown;
