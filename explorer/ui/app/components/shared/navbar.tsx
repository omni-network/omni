import { SearchBar } from './search'
import { ThemeButton } from './themebutton'
import logo from '../../../public/Horizontal_Word Black_Symbol Black.png'
import Logo from './logo'

export default function Navbar() {
  return (
    <header className="static border-none py-3 px-8">
      <div className="w-full flex">
        <div className="px-4 flex gap-12 content-center grid-flow-row">
          <a href="/" className="m-auto">
            <Logo />
          </a>
          <a href="/" className="m-auto text-default text-b font-normal">
            EVM Explorer
          </a>
        </div>

        <div className="grow"></div>
        <div className="flex-none mr-3">
          <ThemeButton />
        </div>
      </div>
    </header>
  )
}
