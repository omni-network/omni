import React from 'react'
interface Props {
  // children: Array<React.ReactNode> | React.ReactNode
  placeholder?: string
}

const SearchBar: React.FC<Props> = ({ placeholder = 'Search', ...props }) => {
  return (
    <div
      className={`relative w-full rounded-[1000px] rounded-l-none bg-search-default bg-opacity-[0.05] border-[1px] border-subtle overflow-hidden`}
    >
      <span className="icon-search absolute top-3 left-3 text-[22px] text-default" />
      <input
        type="text"
        placeholder={placeholder}
        className={`input bg-transparent w-full h-14 px-12 rounded-[1000px] text-subtlest rounded-l-none text-cb`}
      />
    </div>
  )
}

export default SearchBar
