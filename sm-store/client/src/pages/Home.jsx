import { useState, useEffect } from "react";
import {
	ShoppingBag,
	Star,
	ArrowRight,
	Zap,
	Shield,
	Truck,
	RefreshCw,
	ChevronRight,
	Sparkles,
	TrendingUp,
	Heart,
	Play,
} from "lucide-react";

import { Link } from 'react-router-dom';
import { POPPINS } from '@/config/style';

import Navbar from '@/components/Navbar';

import WirelessHeadPhone from '@/assets/wireless-headphone-fb-sm.webp';
import SmartWatch from '@/assets/smart-watch-fb-sm.webp';
import LeatherBag from '@/assets/leather-bag-fb-sm.webp';
import Shoe from '@/assets/shoe-fb-sm.webp';

// ─── Config ─────────────────────────────────────────────────────────────────
const GRADIENT = "linear-gradient(160deg, #4f1899 0%, #7C3AED 30%, #C026D3 62%, #f97316 100%)";
const BLOB_STYLE = "absolute rounded-full pointer-events-none";

// ─── Shared Components ───────────────────────────────────────────────────────
function GradientButton({ children, onClick, className = "", outlined = false }) {
	if (outlined) {
		return (
			<button
				onClick={onClick}
				className={`relative inline-flex items-center gap-2 px-6 py-3 rounded-2xl font-semibold text-sm border-2 border-purple-600 text-purple-600 hover:bg-purple-50 transition-all duration-200 ${className}`}
				style={{ fontFamily: POPPINS }}
			>
				{children}
			</button>
		);
	}
	return (
		<button
			onClick={onClick}
			className={`relative inline-flex items-center gap-2 px-6 py-3.5 rounded-2xl font-semibold text-sm text-white shadow-lg hover:shadow-purple-400/40 hover:-translate-y-0.5 active:scale-95 transition-all duration-200 ${className}`}
			style={{ background: GRADIENT, fontFamily: POPPINS }}
		>
			{children}
		</button>
	);
}

function Badge({ children, color = "purple" }) {
	const colors = {
		purple: "bg-purple-100 text-purple-700",
		orange: "bg-orange-100 text-orange-600",
		pink: "bg-pink-100 text-pink-700",
		green: "bg-green-100 text-green-700",
	};
	return (
		<span className={`inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-semibold tracking-wide ${colors[color]}`}
			style={{ fontFamily: POPPINS }}>
			{children}
		</span>
	);
}



// ─── Hero ─────────────────────────────────────────────────────────────────────
function Hero() {
	return (
		<section className="relative min-h-screen flex items-center overflow-hidden bg-gray-50 pt-20">
			{/* Background blobs */}
			<div className={`${BLOB_STYLE} -top-32 -right-32 w-[700px] h-[700px] opacity-10`}
				style={{ background: "radial-gradient(circle, #7C3AED, transparent 70%)" }} />
			<div className={`${BLOB_STYLE} -bottom-40 -left-40 w-[600px] h-[600px] opacity-10`}
				style={{ background: "radial-gradient(circle, #C026D3, transparent 70%)" }} />
			<div className={`${BLOB_STYLE} top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[400px] h-[400px] opacity-5`}
				style={{ background: "radial-gradient(circle, #f97316, transparent 70%)" }} />

			<div className="relative z-10 max-w-7xl mx-auto px-6 grid lg:grid-cols-2 gap-16 items-center py-24">
				{/* Left copy */}
				<div>
					<div className="flex items-center gap-2 mb-6">
						<Badge color="purple"><TrendingUp size={11} /> Trending Now</Badge>
						<Badge color="orange"><Zap size={11} /> Flash Sale — 40% Off</Badge>
					</div>

					<h1 className="text-5xl lg:text-6xl font-bold text-gray-900 leading-tight tracking-tight mb-6"
						style={{ fontFamily: POPPINS }}>
						Shop Smarter,{" "}
						<span className="relative inline-block">
							<span style={{ background: GRADIENT, WebkitBackgroundClip: "text", WebkitTextFillColor: "transparent" }}>
								Live Brighter
							</span>
						</span>
						.
					</h1>

					<p className="text-base text-gray-500 leading-relaxed mb-10 max-w-md"
						style={{ fontFamily: POPPINS }}>
						Discover curated collections across every category — from everyday essentials to the
						latest trends. Fast shipping, easy returns, and prices that make sense.
					</p>

					<div className="flex flex-wrap gap-4 mb-12">
						<Link to="/shop">
							<GradientButton className="text-base px-8 py-4 h-full">
								Explore Collections <ArrowRight size={16} />
							</GradientButton>
						</Link>
						<Link to="/">
							<GradientButton outlined className="text-base px-8 py-4">
								<Play size={16} /> Watch Story
							</GradientButton>
						</Link>
					</div>

					{/* Social proof strip */}
					<div className="flex items-center gap-8">
						<div>
							<p className="text-2xl font-bold text-gray-900" style={{ fontFamily: POPPINS }}>40K+</p>
							<p className="text-xs text-gray-400">Happy Customers</p>
						</div>
						<div className="w-px h-10 bg-gray-200" />
						<div>
							<div className="flex items-center gap-1 mb-0.5">
								{[...Array(5)].map((_, i) => (
									<Star key={i} size={14} className="fill-yellow-400 text-yellow-400" />
								))}
							</div>
							<p className="text-xs text-gray-400">4.9 / 5 Rating</p>
						</div>
						<div className="w-px h-10 bg-gray-200" />
						<div>
							<p className="text-2xl font-bold text-gray-900" style={{ fontFamily: POPPINS }}>99.9%</p>
							<p className="text-xs text-gray-400">On-Time Delivery</p>
						</div>
					</div>
				</div>

				{/* Right: hero visual card */}
				<div className="relative flex justify-center items-center">
					{/* Gradient backdrop */}
					<div className="absolute inset-0 rounded-3xl opacity-90"
						style={{ background: GRADIENT, transform: "rotate(-3deg) scale(0.95)", borderRadius: "2rem" }} />

					{/* Card */}
					<div className="relative w-full max-w-sm rounded-3xl overflow-hidden shadow-2xl"
						style={{ background: "rgba(255,255,255,0.10)", backdropFilter: "blur(16px)", border: "1px solid rgba(255,255,255,0.2)" }}>
						{/* Mock product display */}
						<div className="p-6 flex flex-col gap-4">
							{/* Top badge */}
							<div className="flex items-center justify-between">
								<span className="px-3 py-1 rounded-full text-xs font-semibold text-white"
									style={{ background: "rgba(255,255,255,0.2)" }}>✦ New Arrivals</span>
								<button className="w-8 h-8 rounded-full flex items-center justify-center"
									style={{ background: "rgba(255,255,255,0.15)" }}>
									<Heart size={14} className="text-white" />
								</button>
							</div>

							{/* SVG Product Illustration */}
							<svg viewBox="0 0 340 200" fill="none" className="w-full" xmlns="http://www.w3.org/2000/svg">
								{/* Main card */}
								<rect x="40" y="20" width="260" height="160" rx="20" fill="rgba(255,255,255,0.08)" stroke="rgba(255,255,255,0.2)" strokeWidth="1.5" />
								{/* Shelf line */}
								<rect x="60" y="130" width="220" height="3" rx="2" fill="rgba(255,255,255,0.15)" />
								{/* Product 1 - bag shape */}
								<rect x="75" y="75" width="55" height="55" rx="10" fill="rgba(124,58,237,0.6)" stroke="rgba(255,255,255,0.25)" strokeWidth="1" />
								<path d="M90 75 Q102 60 115 75" stroke="rgba(255,255,255,0.5)" strokeWidth="2" fill="none" strokeLinecap="round" />
								<circle cx="102" cy="105" r="5" fill="rgba(255,255,255,0.3)" />
								{/* Product 2 - sneaker */}
								<rect x="145" y="70" width="65" height="60" rx="10" fill="rgba(192,38,211,0.5)" stroke="rgba(255,255,255,0.2)" strokeWidth="1" />
								<ellipse cx="177" cy="115" rx="22" ry="8" fill="rgba(255,255,255,0.15)" />
								<path d="M155 100 Q177 80 205 95" stroke="rgba(255,255,255,0.5)" strokeWidth="1.5" fill="none" />
								{/* Product 3 - watch */}
								<rect x="228" y="75" width="52" height="55" rx="10" fill="rgba(249,115,22,0.5)" stroke="rgba(255,255,255,0.2)" strokeWidth="1" />
								<circle cx="254" cy="102" r="14" fill="rgba(255,255,255,0.12)" stroke="rgba(255,255,255,0.3)" strokeWidth="1.5" />
								<line x1="254" y1="96" x2="254" y2="102" stroke="rgba(255,255,255,0.7)" strokeWidth="1.5" strokeLinecap="round" />
								<line x1="254" y1="102" x2="259" y2="107" stroke="rgba(255,255,255,0.7)" strokeWidth="1.5" strokeLinecap="round" />
								{/* Floating badge - star */}
								<circle cx="60" cy="42" r="14" fill="rgba(249,115,22,0.7)" stroke="rgba(255,255,255,0.3)" strokeWidth="1" />
								<text x="60" y="47" textAnchor="middle" fontSize="12" fill="white">★</text>
								{/* Floating badge - cart */}
								<circle cx="290" cy="160" r="16" fill="rgba(74,222,128,0.7)" stroke="rgba(255,255,255,0.3)" strokeWidth="1" />
								<text x="290" y="165" textAnchor="middle" fontSize="12" fill="white">+</text>
								{/* Sparkle dots */}
								<circle cx="165" cy="38" r="3" fill="rgba(255,255,255,0.5)" />
								<circle cx="220" cy="55" r="2" fill="rgba(255,255,255,0.4)" />
								<circle cx="85" cy="165" r="2.5" fill="rgba(255,255,255,0.4)" />
							</svg>

							{/* Product info */}
							<div className="space-y-1">
								<p className="text-xs text-white/60 uppercase tracking-widest">Featured Drop</p>
								<p className="text-base font-bold text-white">Summer Essentials Bundle</p>
								<div className="flex items-center justify-between">
									<span className="text-lg font-bold text-white">₱2,499</span>
									<span className="text-xs line-through text-white/40">₱4,199</span>
								</div>
							</div>

							{/* CTA inside card */}
							<button className="w-full py-3 rounded-2xl text-sm font-semibold text-white transition-all hover:opacity-90"
								style={{ background: "rgba(255,255,255,0.18)", border: "1px solid rgba(255,255,255,0.25)", fontFamily: POPPINS }}>
								Add to Cart →
							</button>
						</div>
					</div>

					{/* Floating tag card */}
					<div className="absolute -bottom-4 -left-8 bg-white rounded-2xl shadow-xl px-4 py-3 flex items-center gap-3"
						style={{ fontFamily: POPPINS }}>
						<div className="w-8 h-8 rounded-xl flex items-center justify-center"
							style={{ background: "linear-gradient(135deg, #22c55e, #16a34a)" }}>
							<Truck size={14} className="text-white" />
						</div>
						<div>
							<p className="text-xs font-bold text-gray-800">Free Shipping</p>
							<p className="text-xs text-gray-400">Orders over ₱999</p>
						</div>
					</div>

					{/* Floating review card */}
					<div className="absolute -top-4 -right-4 bg-white rounded-2xl shadow-xl px-4 py-3"
						style={{ fontFamily: POPPINS }}>
						<div className="flex items-center gap-1 mb-1">
							{[...Array(5)].map((_, i) => <Star key={i} size={10} className="fill-yellow-400 text-yellow-400" />)}
						</div>
						<p className="text-xs font-semibold text-gray-800">"Absolutely love it!"</p>
						<p className="text-xs text-gray-400">— Maria G.</p>
					</div>
				</div>
			</div>
		</section>
	);
}

// ─── Category Strip ───────────────────────────────────────────────────────────
function Categories() {
	const cats = [
		{ name: "Electronics", icon: "⚡", color: "from-violet-500 to-purple-700" },
		{ name: "Fashion", icon: "👗", color: "from-pink-500 to-rose-600" },
		{ name: "Home & Living", icon: "🏠", color: "from-orange-400 to-amber-600" },
		{ name: "Sports", icon: "🏋️", color: "from-green-400 to-emerald-600" },
		{ name: "Beauty", icon: "💄", color: "from-fuchsia-500 to-pink-600" },
		{ name: "Books", icon: "📚", color: "from-sky-400 to-blue-600" },
	];

	return (
		<section className="py-20 bg-white relative overflow-hidden" style={{ fontFamily: POPPINS }}>
			{/* Blobs */}
			<div className={`${BLOB_STYLE} top-0 right-0 w-96 h-96 opacity-5`}
				style={{ background: "radial-gradient(circle, #7C3AED, transparent 70%)" }} />

			<div className="max-w-7xl mx-auto px-6">
				<div className="flex items-end justify-between mb-10">
					<div>
						<Badge color="purple"><Sparkles size={11} /> Categories</Badge>
						<h2 className="text-3xl font-bold text-gray-900 mt-3 tracking-tight">Shop by Category</h2>
						<p className="text-sm text-gray-400 mt-1">Everything you need, all in one place.</p>
					</div>
					<a href="#" className="hidden md:flex items-center gap-1 text-sm font-semibold text-purple-600 hover:text-purple-700 transition-colors">
						View All <ChevronRight size={16} />
					</a>
				</div>

				<div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-4">
					{cats.map((cat) => (
						<button 
							key={cat.name}
							className="group relative flex flex-col items-center gap-3 p-6 rounded-2xl bg-gray-50 hover:shadow-lg hover:-translate-y-1 transition-all duration-200 border border-gray-100"
						>
							<div className={`w-14 h-14 rounded-2xl bg-gradient-to-br ${cat.color} flex items-center justify-center text-2xl shadow-md group-hover:scale-110 transition-transform duration-200`}>
								{cat.icon}
							</div>
							<span className="text-xs font-semibold text-gray-700 text-center">{cat.name}</span>
						</button>
					))}
				</div>
			</div>
		</section>
	);
}

// ─── Product Card ─────────────────────────────────────────────────────────────
function ProductCard({ name, price, originalPrice, rating, reviews, badge, img, gradient }) {
	const [wished, setWished] = useState(false);
	return (
		<div className="group relative bg-white rounded-3xl overflow-hidden border border-gray-100 hover:shadow-xl hover:-translate-y-1 transition-all duration-300 shadow"
			style={{ fontFamily: POPPINS }}>
			{/* Image area */}
			<div className={`relative h-52 bg-gradient-to-br ${gradient} flex items-center justify-center`}>
				<img src={img} alt={name} className="h-52 object-cover" />

				{badge && (
					<div className="absolute top-3 left-3">
						<span className="px-2.5 py-1 rounded-full text-xs font-bold text-white"
							style={{ background: "rgba(0,0,0,0.3)", backdropFilter: "blur(4px)" }}>
							{badge}
						</span>
					</div>
				)}

				<button onClick={() => setWished(!wished)}
					className="absolute top-3 right-3 w-8 h-8 rounded-full bg-white/80 backdrop-blur-sm flex items-center justify-center shadow hover:scale-110 transition-transform">
					<Heart size={14} className={wished ? "fill-red-500 text-red-500" : "text-gray-400"} />
				</button>

				{/* Quick add hover */}
				<div className="absolute inset-x-3 bottom-3 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
					<button className="w-full py-2 rounded-xl text-xs font-bold text-white shadow-lg"
						style={{ background: GRADIENT }}>
						Quick Add +
					</button>
				</div>
			</div>

			{/* Info */}
			<div className="p-4">
				<p className="text-sm font-semibold text-gray-800 truncate">{name}</p>
				<div className="flex items-center gap-1 mt-1">
					<Star size={12} className="fill-yellow-400 text-yellow-400" />
					<span className="text-xs font-medium text-gray-700">{rating}</span>
					<span className="text-xs text-gray-400">({reviews})</span>
				</div>
				<div className="flex items-center justify-between mt-3">
					<div>
						<span className="text-base font-bold text-gray-900">₱{price.toLocaleString()}</span>
						{originalPrice && (
							<span className="text-xs text-gray-400 line-through ml-2">₱{originalPrice.toLocaleString()}</span>
						)}
					</div>
					<button className="w-8 h-8 rounded-xl flex items-center justify-center text-white shadow"
						style={{ background: GRADIENT }}>
						<ShoppingBag size={14} />
					</button>
				</div>
			</div>
		</div>
	);
}

// ─── Featured Products ────────────────────────────────────────────────────────
function FeaturedProducts() {
	const products = [
		{ 
			name: "Wireless Noise-Cancelling Headphones", 
			price: 1200, 
			originalPrice: 1300, 
			rating: 4.9, 
			reviews: 1204, 
			badge: "Best Seller", 
			img: WirelessHeadPhone,
			gradient: "from-violet-100 to-purple-200" 
		},
		{ 
			name: "Ultra-Slim Smartwatch Pro", 
			price: 7999, 
			originalPrice: 9739, 
			rating: 4.8, 
			reviews: 876, 
			badge: "40% Off", 
			img: SmartWatch,
			gradient: "from-slate-100 to-gray-200" 
		},
		{ 
			name: "Leather Tote Bag — Caramel", 
			price: 2199, 
			originalPrice: null, 
			rating: 4.7, 
			reviews: 532, 
			badge: "New", 
			img: LeatherBag, 
			gradient: "from-amber-100 to-orange-200" 
		},
		{ 
			name: "Running Sneakers Air Max", 
			price: 2399, 
			originalPrice: 4599, 
			rating: 4.9, 
			reviews: 2140, 
			badge: "Hot", 
			img: Shoe,
			gradient: "from-rose-100 to-pink-200" 
		},
	];

	return (
		<section className="py-20 bg-gray-50 relative overflow-hidden" style={{ fontFamily: POPPINS }}>
			<div className={`${BLOB_STYLE} -bottom-20 left-1/2 w-[500px] h-[500px] opacity-5 -translate-x-1/2`}
				style={{ background: "radial-gradient(circle, #C026D3, transparent 70%)" }} />

			<div className="max-w-7xl mx-auto px-6">
				<div className="flex items-end justify-between mb-10">
					<div>
						<Badge color="pink"><TrendingUp size={11} /> Top Picks</Badge>
						<h2 className="text-3xl font-bold text-gray-900 mt-3 tracking-tight">Featured Products</h2>
						<p className="text-sm text-gray-400 mt-1">Hand-picked items just for you.</p>
					</div>
					<a href="#" className="hidden md:flex items-center gap-1 text-sm font-semibold text-purple-600 hover:text-purple-700 transition-colors">
						See All <ChevronRight size={16} />
					</a>
				</div>

				<div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
					{products.map((p) => <ProductCard key={p.name} {...p} />)}
				</div>
			</div>
		</section>
	);
}

// ─── Promo Banner ─────────────────────────────────────────────────────────────
function PromoBanner() {
	return (
		<section className="py-16 px-6" style={{ fontFamily: POPPINS }}>
			<div className="max-w-7xl mx-auto rounded-3xl overflow-hidden relative" style={{ background: GRADIENT }}>
				{/* Blobs */}
				<div className={`${BLOB_STYLE} -top-16 -right-16 w-72 h-72`}
					style={{ background: "rgba(255,255,255,0.07)" }} />
				<div className={`${BLOB_STYLE} -bottom-20 -left-20 w-80 h-80`}
					style={{ background: "rgba(255,150,50,0.12)" }} />

				<div className="relative z-10 grid lg:grid-cols-2 items-center gap-8 px-10 py-14">
					<div>
						<div className="inline-flex items-center gap-2 px-4 py-1.5 rounded-full text-xs font-medium text-white/90 border border-white/20 mb-6"
							style={{ background: "rgba(255,255,255,0.12)", backdropFilter: "blur(4px)" }}>
							<span className="w-1.5 h-1.5 rounded-full bg-green-400" />
							Limited Time Offer
						</div>
						<h2 className="text-3xl lg:text-4xl font-bold text-white leading-tight mb-4">
							Summer Sale is Live.<br />Up to 60% Off.
						</h2>
						<p className="text-white/70 text-sm leading-relaxed mb-8 max-w-sm">
							Shop hundreds of deals across all categories. Free delivery on orders above ₱999. No code needed — discount applied automatically.
						</p>
						<div className="flex gap-4">
							<button className="px-8 py-3.5 rounded-2xl text-sm font-bold text-purple-700 bg-white hover:bg-gray-50 shadow-lg hover:-translate-y-0.5 active:scale-95 transition-all">
								Shop the Sale →
							</button>
							<button className="px-8 py-3.5 rounded-2xl text-sm font-bold text-white border-2 border-white/30 hover:bg-white/10 transition-all">
								Learn More
							</button>
						</div>
					</div>

					{/* SVG illustration */}
					<div className="flex justify-center lg:justify-end">
						<svg viewBox="0 0 320 200" fill="none" className="w-full max-w-xs" xmlns="http://www.w3.org/2000/svg">
							{/* Tag shape */}
							<rect x="60" y="30" width="200" height="140" rx="20" fill="rgba(255,255,255,0.08)" stroke="rgba(255,255,255,0.2)" strokeWidth="1.5" />
							{/* Percent sign area */}
							<circle cx="120" cy="85" r="25" fill="rgba(124,58,237,0.6)" stroke="rgba(255,255,255,0.3)" strokeWidth="1.5" />
							<circle cx="108" cy="74" r="6" fill="rgba(255,255,255,0.4)" />
							<circle cx="132" cy="96" r="6" fill="rgba(255,255,255,0.4)" />
							<line x1="110" y1="96" x2="130" y2="74" stroke="rgba(255,255,255,0.7)" strokeWidth="2.5" strokeLinecap="round" />
							{/* Price tag */}
							<rect x="165" y="62" width="70" height="50" rx="12" fill="rgba(249,115,22,0.6)" stroke="rgba(255,255,255,0.25)" strokeWidth="1" />
							<circle cx="178" cy="80" r="5" fill="rgba(255,255,255,0.4)" />
							<line x1="190" y1="78" x2="222" y2="78" stroke="rgba(255,255,255,0.4)" strokeWidth="1.5" strokeLinecap="round" />
							<line x1="190" y1="86" x2="218" y2="86" stroke="rgba(255,255,255,0.25)" strokeWidth="1.5" strokeLinecap="round" />
							{/* Floating badges */}
							<circle cx="68" cy="50" r="18" fill="rgba(74,222,128,0.7)" stroke="rgba(255,255,255,0.3)" strokeWidth="1" />
							<text x="68" y="55" textAnchor="middle" fontSize="13" fill="white" fontWeight="bold">✓</text>
							<circle cx="254" cy="160" r="18" fill="rgba(192,38,211,0.7)" stroke="rgba(255,255,255,0.3)" strokeWidth="1" />
							<text x="254" y="165" textAnchor="middle" fontSize="11" fill="white" fontWeight="bold">60%</text>
							{/* Sparks */}
							<circle cx="90" cy="160" r="3" fill="rgba(255,255,255,0.5)" />
							<circle cx="250" cy="40" r="2.5" fill="rgba(255,255,255,0.4)" />
							<circle cx="200" cy="155" r="2" fill="rgba(255,255,255,0.3)" />
						</svg>
					</div>
				</div>
			</div>
		</section>
	);
}

// ─── Trust Strip ──────────────────────────────────────────────────────────────
function TrustStrip() {
	const perks = [
		{ icon: <Truck size={22} />, title: "Free Shipping", desc: "On orders over ₱999", color: "text-purple-600 bg-purple-50" },
		{ icon: <RefreshCw size={22} />, title: "Easy Returns", desc: "30-day hassle-free returns", color: "text-pink-600 bg-pink-50" },
		{ icon: <Shield size={22} />, title: "Secure Payment", desc: "256-bit SSL encryption", color: "text-orange-600 bg-orange-50" },
		{ icon: <Zap size={22} />, title: "Fast Delivery", desc: "Same-day in Metro Manila", color: "text-green-600 bg-green-50" },
	];

	return (
		<section className="py-16 bg-white" style={{ fontFamily: POPPINS }}>
			<div className="max-w-7xl mx-auto px-6 grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
				{perks.map((p) => (
					<div key={p.title} className="flex items-start gap-4 p-6 rounded-2xl bg-gray-50 border border-gray-100">
						<div className={`w-12 h-12 rounded-2xl flex items-center justify-center flex-shrink-0 ${p.color}`}>
							{p.icon}
						</div>
						<div>
							<p className="text-sm font-bold text-gray-800">{p.title}</p>
							<p className="text-xs text-gray-400 mt-0.5">{p.desc}</p>
						</div>
					</div>
				))}
			</div>
		</section>
	);
}

// ─── Testimonials ─────────────────────────────────────────────────────────────
function Testimonials() {
	const reviews = [
		{ name: "Maria G.", role: "Verified Buyer", text: "Fast shipping and the product quality exceeded my expectations. Luminary has become my go-to shopping app!", avatar: "👩‍🦰", rating: 5 },
		{ name: "Carlo R.", role: "Verified Buyer", text: "The summer sale deals are insane. Got a smartwatch for half the price and it arrived in perfect condition.", avatar: "👨‍💼", rating: 5 },
		{ name: "Jana L.", role: "Verified Buyer", text: "Customer support is super responsive. Returns were so easy. Already ordered 5 times this month 😂", avatar: "👩‍💻", rating: 5 },
	];

	return (
		<section className="py-20 bg-gray-50 relative overflow-hidden" style={{ fontFamily: POPPINS }}>
			<div className={`${BLOB_STYLE} -top-10 right-0 w-96 h-96 opacity-5`}
				style={{ background: "radial-gradient(circle, #7C3AED, transparent 70%)" }} />

			<div className="max-w-7xl mx-auto px-6">
				<div className="text-center mb-12">
					<Badge color="green"><Star size={11} className="fill-current" /> Reviews</Badge>
					<h2 className="text-3xl font-bold text-gray-900 mt-3 tracking-tight">What Our Customers Say</h2>
					<p className="text-sm text-gray-400 mt-1">Real reviews from real shoppers.</p>
				</div>

				<div className="grid md:grid-cols-3 gap-6">
					{reviews.map((r) => (
						<div key={r.name} className="bg-white rounded-3xl p-6 border border-gray-100 shadow-sm hover:shadow-md transition-shadow">
							<div className="flex items-center gap-1 mb-4">
								{[...Array(r.rating)].map((_, i) => (
									<Star key={i} size={14} className="fill-yellow-400 text-yellow-400" />
								))}
							</div>
							<p className="text-sm text-gray-600 leading-relaxed mb-5">"{r.text}"</p>
							<div className="flex items-center gap-3 pt-4 border-t border-gray-100">
								<div className="w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center text-xl">{r.avatar}</div>
								<div>
									<p className="text-sm font-bold text-gray-800">{r.name}</p>
									<p className="text-xs text-gray-400">{r.role}</p>
								</div>
							</div>
						</div>
					))}
				</div>
			</div>
		</section>
	);
}

// ─── Newsletter ───────────────────────────────────────────────────────────────
function Newsletter() {
	const [email, setEmail] = useState("");
	const [sent, setSent] = useState(false);

	return (
		<section className="py-20 bg-white relative overflow-hidden" style={{ fontFamily: POPPINS }}>
			<div className={`${BLOB_STYLE} -bottom-16 -left-16 w-72 h-72 opacity-5`}
				style={{ background: "radial-gradient(circle, #C026D3, transparent 70%)" }} />
			<div className={`${BLOB_STYLE} -top-16 -right-16 w-72 h-72 opacity-5`}
				style={{ background: "radial-gradient(circle, #f97316, transparent 70%)" }} />

			<div className="max-w-2xl mx-auto px-6 text-center">
				<Badge color="orange"><Zap size={11} /> Stay in the Loop</Badge>
				<h2 className="text-3xl font-bold text-gray-900 mt-4 mb-2 tracking-tight">Get Exclusive Deals First</h2>
				<p className="text-sm text-gray-400 mb-8">Subscribe to our newsletter and be the first to know about sales, new arrivals, and member-only discounts.</p>

				{sent ? (
					<div className="flex flex-col items-center gap-2">
						<div className="w-14 h-14 rounded-full bg-green-100 flex items-center justify-center text-2xl">✓</div>
						<p className="text-sm font-semibold text-green-600">You're in! Check your inbox soon.</p>
					</div>
				) : (
					<div className="flex gap-3 max-w-md mx-auto">
						<input
							type="email"
							value={email}
							onChange={(e) => setEmail(e.target.value)}
							placeholder="Enter your email"
							className="flex-1 px-4 py-3.5 rounded-2xl border border-gray-200 text-sm text-gray-700 focus:outline-none focus:border-purple-400 placeholder-gray-300"
							style={{ fontFamily: POPPINS }}
						/>
						<GradientButton onClick={() => email && setSent(true)}>
							Subscribe
						</GradientButton>
					</div>
				)}

				<p className="text-xs text-gray-400 mt-4">No spam, ever. Unsubscribe anytime.</p>
			</div>
		</section>
	);
}

// ─── Footer ───────────────────────────────────────────────────────────────────
function Footer() {
	const copyrightLinks = {
		"Privacy Policy": "/privpolicy", 
		"Terms of Service": "/termsofservice", 
		"Cookie Policy": "/cukiepolicy",
	};

	return (
		<footer className="bg-gray-900 text-white pt-16 pb-8" style={{ fontFamily: POPPINS }}>
			<div className="max-w-7xl mx-auto px-6 grid sm:grid-cols-2 lg:grid-cols-4 gap-10 mb-12">
				{/* Brand */}
				<div>
					<div className="flex items-center gap-2 mb-4">
						<div className="w-8 h-8 rounded-xl flex items-center justify-center" style={{ background: GRADIENT }}>
							<Sparkles size={14} className="text-white" />
						</div>
						<span className="text-base font-bold tracking-tight">Luminary</span>
					</div>
					<p className="text-sm text-gray-400 leading-relaxed">Shop smarter. Live brighter. Your one-stop destination for everything you love.</p>
				</div>

				{/* Links */}
				{[
					{ title: "Shop", links: ["New Arrivals", "Best Sellers", "Sale", "Brands"] },
					{ title: "Support", links: ["Help Center", "Track Order", "Returns", "Contact Us"] },
					{ title: "Company", links: ["About", "Careers", "Blog", "Press"] },
				].map((col) => (
					<div key={col.title}>
						<p className="text-xs font-semibold uppercase tracking-widest text-gray-500 mb-4">{col.title}</p>
						<ul className="space-y-2">
							{col.links.map((l) => (
								<li key={l}><a href="#" className="text-sm text-gray-400 hover:text-white transition-colors">{l}</a></li>
							))}
						</ul>
					</div>
				))}
			</div>

			<div className="max-w-7xl mx-auto px-6 pt-8 border-t border-gray-800 flex flex-col sm:flex-row items-center justify-between gap-4">
				<p className="text-xs text-gray-500">© {new Date().getFullYear()} Luminary. All rights reserved.</p>
				<div className="flex gap-6">
					{Object.entries(copyrightLinks).map((l) => {
						console.log(l[1]);
						return (
						<a 
							key={l[0]} 
							href={l[1]} 
							className="text-xs text-gray-500 hover:text-white transition-colors"
						>
							{l[0]}
						</a>
					)})}
				</div>
			</div>
		</footer>
	);
}

// ─── Page ─────────────────────────────────────────────────────────────────────
export default function HomePage() {
	return (
		<>
			<link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600;700;800&display=swap" rel="stylesheet" />
			<div style={{ fontFamily: POPPINS }}>
				<Navbar />
				<Hero />
				<Categories />
				<FeaturedProducts />
				<PromoBanner />
				<TrustStrip />
				<Testimonials />
				<Newsletter />
				<Footer />
			</div>
		</>
	);
}