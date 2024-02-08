import { SearchBar } from './search'
import { ThemeButton } from './themebutton'
import logo from '~/components/ui/Horizontal_Word Black_Symbol Black.png'

export default function Navbar() {
  return (
    <header className="static m-auto border-b">
      <div className="w-full flex max-w-screen-xl m-auto">
        <div className="mx-auto px-4 flex-none grid justify-start">
          <a href="#" className="m-auto">
            <img className="h-8 w-auto" src={logo} alt="" />
          </a>
        </div>
        <div className="flex-auto place-content-stretch">
          <SearchBar />
        </div>
        <div className="grow"></div>
        <div className="flex-none mr-3">
          <ThemeButton />
        </div>
      </div>
    </header>
  )
}
