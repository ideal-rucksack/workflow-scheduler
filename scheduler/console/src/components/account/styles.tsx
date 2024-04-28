import styled from "styled-components";
import {signinRightBg} from "@/assets";

export const Styles = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    width: 100%;
    height: 100vh; // Make sure the parent takes the full viewport height
    .content-container {
        padding: 0 20px;
        display: flex;
        justify-content: center;
        height: 60vh;
        width: 80vw;
        .signin-wrapper {
            width: 40%;
            display: flex;
            justify-content: space-between;
            flex-direction: column;
        }
        .signin-container {
            .title-container {
                display: flex;
                justify-content: center;
            }
            .signin-form-container {
                .signin-form {
                    .signin-buttons {
                        .signin-button {
                            width: 100%;
                        }
                    }
                }
            }
            .to-signup {
                display: flex;
                justify-content: center;
            }
            .signin-right {
                background-image: url(${signinRightBg});
                background-size: contain;
                background-position: center;
                background-repeat: no-repeat;
                width: 100%;
                height: 100%;
            }
        }
        .right-wrapper {
            width: 60%;
        }
    }
    
`;
