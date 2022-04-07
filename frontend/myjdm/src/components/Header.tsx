
import logo from '../images/logo.webp'
import { HeaderComponent } from '../styles/styles.js'

export function Header() {
  return (
      <HeaderComponent>
        <div>
            <img src={logo} className="App-logo" alt="logo" />
            <ul>
                <li>Cadastrar</li>
                <li>Veículos</li>
                <li>Contatos</li>
            </ul>
        </div>
      </HeaderComponent>
  )
}
