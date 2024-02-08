import { SocialIcon } from 'react-social-icons/component'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faDiscord, faXTwitter, faTelegram } from '@fortawesome/free-brands-svg-icons'
import 'react-social-icons/discord'
import 'react-social-icons/twitter'
import 'react-social-icons/telegram'
import logo from '~/components/ui/Horizontal_Word Black_Symbol Black.png'

export function Footer() {
  return (
    <footer className="mt-auto border-t">
      <div className="w-full max-w-screen-xl mx-auto grid grid-rows-2 grid-cols-3 gap-9 px-4 py-10">
        {/* top left */}
        <section className="place-items-start">
          <a href="#" className="-m-1.5 p-1.5">
            <span className="sr-only">Omni Network</span>
            <img className="h-8 w-auto" src={logo} alt="" />
          </a>
        </section>
        {/* top middle */}
        <section className="grid grid-flow-col gap-6 place-items-center">
          <a className="m3 prose prose-m" href="https://omni.network/">
            Omni Network
          </a>

          <a className="m3 prose prose-m" href="https://docs.omni.network/">
            Docs
          </a>
          <a className="m3 prose prose-m" href="https://explorer.omni.network">
            Explorer
          </a>
          <a className="m3 prose prose-m" href="https://news.omni.network/">
            Blog
          </a>
        </section>
        {/* top right */}
        <section className=""></section>
        {/* bottom left */}
        <section className="place-items-center m3">
          <a className="prose prose-sm">© 2023 The Omni Network | All Rights Reserved.</a>
        </section>
        {/* bottom middle */}
        <section className="m3"></section>
        {/* bottom right */}
        <section className="grid grid-flow-col gap-6 place-items-center-stretch">
          <a href="https://discord.gg/bKNXmaX9VD" target="_blank" className="m3 place-items-center">
            <FontAwesomeIcon icon={faDiscord} size="xl" />
          </a>
          <a href="https://twitter.com/OmniFDN" target="_blank" className="m3">
            <FontAwesomeIcon icon={faXTwitter} size="xl" />
          </a>
          <a href="https://t.me/OmniFDN" target="_blank" className="m3">
            <FontAwesomeIcon icon={faTelegram} size="xl" className="m3" />
          </a>
        </section>
      </div>
    </footer>
  )
}
