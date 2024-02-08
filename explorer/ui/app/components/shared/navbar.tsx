import { SearchBar } from './search'
import { ThemeButton } from './themebutton'
import logo from '~/components/ui/Horizontal_Word Black_Symbol Black.png'

export default function Navbar() {
  return (
    <nav className="sticky top-0 z-50">
      <div className="navbar w-full shadow-md">
        <div className="flex-1">
          <a href="#" className="m3 p-1.5">
            <img className="h-8 w-auto" src={logo} alt="" />
          </a>
          <SearchBar />
        </div>
        <div className="flex-none gap-2 ml-3 mr-3">
          <ThemeButton />
        </div>
      </div>
    </nav>
  )
}
