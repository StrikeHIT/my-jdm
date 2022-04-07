import styled from 'styled-components';

export const HeaderComponent = styled.header`
    padding:20px 0px;
    background-color: #000;
    > div{
        max-width: 1300px;
        margin: 0px auto;
        display: flex;
        justify-content: space-between;
        align-items: center;
        img{
            width: 190px;
            height: 80px;
            object-fit: cover;
        }
        ul{
            display: flex;
            justify-content: space-between;
            max-width: 300px;
            width: 95%;
            list-style: none;
            li{
                color: #FFF;
            }
        }
    }
`; 

export const FooterComponent = styled.header`
    padding:15px 0px 20px 0px;
    background-color: #000;
    margin-top: 30px;
    > p{
        color: #FFF;
        text-align: center;
        font-weight: 300;
        margin: 0px;
    }
`; 