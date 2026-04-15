

const IllustrationSVG = () => (
  <svg viewBox="0 0 340 200" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-full max-w-sm">
    {/* Product card */}
    <rect x="20" y="60" width="300" height="130" rx="18"
      fill="rgba(255,255,255,0.08)" stroke="rgba(255,255,255,0.18)" strokeWidth="1" />

    {/* Product image placeholder */}
    <rect x="36" y="76" width="56" height="56" rx="10" fill="rgba(255,255,255,0.12)" stroke="rgba(255,255,255,0.15)" strokeWidth="1" />
    <rect x="50" y="90" width="28" height="20" rx="4" fill="rgba(255,200,120,0.4)" />
    <circle cx="64" cy="94" r="5" fill="rgba(255,200,120,0.6)" />

    {/* Product name + category */}
    <rect x="104" y="80" width="100" height="9" rx="4.5" fill="rgba(255,255,255,0.25)" />
    <rect x="104" y="97" width="64" height="7" rx="3.5" fill="rgba(255,255,255,0.12)" />

    {/* Price */}
    <rect x="104" y="114" width="56" height="10" rx="5" fill="rgba(255,200,120,0.45)" />

    {/* Divider */}
    <rect x="36" y="144" width="268" height="1" rx="0.5" fill="rgba(255,255,255,0.08)" />

    {/* Rating dots */}
    {[0,1,2,3,4].map(i => (
      <rect key={i} x={36 + i * 16} y="156" width="10" height="10" rx="5"
        fill={i < 4 ? "rgba(255,200,120,0.6)" : "rgba(255,255,255,0.15)"} />
    ))}
    <rect x="120" y="159" width="40" height="5" rx="2.5" fill="rgba(255,255,255,0.15)" />

    {/* Add to cart button */}
    <rect x="222" y="152" width="86" height="28" rx="9"
      fill="rgba(255,255,255,0.15)" stroke="rgba(255,255,255,0.25)" strokeWidth="1" />
    <rect x="234" y="162" width="62" height="7" rx="3.5" fill="rgba(255,255,255,0.6)" />

    {/* Floating cart badge */}
    <circle cx="262" cy="36" r="24" fill="rgba(255,255,255,0.1)" stroke="rgba(255,255,255,0.2)" strokeWidth="1" />
    <circle cx="262" cy="36" r="14" fill="rgba(249,115,22,0.5)" />
    {/* Cart icon */}
    <path d="M255 31h2l2 8h8l1.5-5h-10" stroke="white" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
    <circle cx="261" cy="41" r="1" fill="white" />
    <circle cx="265" cy="41" r="1" fill="white" />

    {/* Floating discount badge */}
    <circle cx="94" cy="18" r="16" fill="rgba(255,255,255,0.1)" stroke="rgba(255,255,255,0.2)" strokeWidth="1" />
    <rect x="83" y="11" width="22" height="14" rx="4" fill="rgba(192,38,211,0.6)" />
    <rect x="87" y="15" width="14" height="5" rx="2.5" fill="rgba(255,255,255,0.8)" />

    {/* Floating delivery badge */}
    <circle cx="180" cy="8" r="10" fill="rgba(255,255,255,0.1)" stroke="rgba(255,255,255,0.2)" strokeWidth="1" />
    <circle cx="180" cy="8" r="5" fill="rgba(96,165,250,0.7)" />
    {/* Tiny truck shape */}
    <rect x="175" y="6" width="7" height="4" rx="1" fill="white" opacity="0.9" />
    <rect x="180" y="7" width="3" height="3" rx="0.5" fill="rgba(96,165,250,0.9)" />
  </svg>
);

export default IllustrationSVG;