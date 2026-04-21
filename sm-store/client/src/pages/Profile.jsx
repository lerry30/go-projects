import { useState, useEffect } from "react";
import {
	User, MapPin, CreditCard, ShoppingBag, Settings,
	Edit3, Plus, Trash2, Check, ChevronRight, Package,
	Star, LogOut, Camera, Shield, Bell, Lock, ChevronDown
} from "lucide-react";

import { GradientButton } from '@/components/Buttons';
import { PrimaryField } from '@/components/Fields'
import { ErrorModal } from '@/components/Modal';

import { poppins, gradient } from '@/config/style';

import { apiRequest, sendJSON } from '@/utils/send';
import { formatDate } from '@/utils/datetime';
import { BASE_URL } from '@/config/server';

import { useNavigate } from 'react-router-dom';

import Loading from '@/components/Loading'

const panelGradient = "linear-gradient(160deg, #4f1899 0%, #7C3AED 30%, #C026D3 62%, #f97316 100%)";

const TABS = [
	{ id: "profile", label: "Profile", icon: User },
	{ id: "address", label: "Address", icon: MapPin },
	{ id: "payment", label: "Payment", icon: CreditCard },
	{ id: "orders", label: "Orders", icon: ShoppingBag },
	{ id: "settings", label: "Settings", icon: Settings },
];

const mockOrders = [
	{ id: "#LM-8821", date: "Apr 10, 2025", status: "Delivered", total: "$129.00", items: 3, rating: 5 },
	{ id: "#LM-7743", date: "Mar 28, 2025", status: "In Transit", total: "$74.50", items: 1, rating: null },
	{ id: "#LM-6610", date: "Mar 15, 2025", status: "Delivered", total: "$210.00", items: 2, rating: 4 },
	{ id: "#LM-5502", date: "Feb 22, 2025", status: "Delivered", total: "$55.00", items: 1, rating: 5 },
];

const statusColor = {
	"Delivered": "text-green-600 bg-green-50",
	"In Transit": "text-orange-500 bg-orange-50",
	"Cancelled": "text-red-500 bg-red-50",
};

// ── Logo ──────────────────────────────────────────────────────────────────────
function Logo() {
	return (
		<div className="flex items-center gap-2 mb-10">
			<div className="w-8 h-8 rounded-xl flex items-center justify-center"
				style={{ background: gradient }}>
				<svg width="14" height="14" viewBox="0 0 14 14" fill="none">
					<path d="M7 1L9 5.5H13.5L9.8 8.5L11.2 13L7 10.2L2.8 13L4.2 8.5L0.5 5.5H5L7 1Z"
						fill="white" />
				</svg>
			</div>
			<span className="text-base font-semibold text-gray-900 tracking-tight" style={{ fontFamily: poppins }}>Luminary</span>
		</div>
	);
}

// ── Toast ─────────────────────────────────────────────────────────────────────
function Toast({ msg }) {
	if (!msg) return null;
	return (
		<div className="fixed bottom-6 left-1/2 -translate-x-1/2 z-50 flex items-center gap-2 bg-gray-900 text-white text-xs font-medium px-5 py-3 rounded-2xl shadow-xl"
			style={{ fontFamily: poppins }}>
			<Check size={14} className="text-green-400" /> {msg}
		</div>
	);
}

// ── RIGHT PANEL ───────────────────────────────────────────────────────────────
function RightPanel({ activeTab }) {
	const content = {
		profile: {
			badge: "Your public profile",
			headline: <>Build your<br />Luminary identity.</>,
			sub: "Keep your info accurate so your orders and communications always reach you.",
		},
		address: {
			badge: "Saved addresses",
			headline: <>Ship anywhere<br />without friction.</>,
			sub: "Store multiple delivery addresses so checkout is always one tap away.",
		},
		payment: {
			badge: "Secure payments",
			headline: <>Pay your way,<br />every time.</>,
			sub: "Your card data is encrypted and never stored on our servers.",
		},
		orders: {
			badge: "Order history",
			headline: <>Everything you've<br />ever loved.</>,
			sub: "Revisit past purchases, track packages, and re-order favourites in seconds.",
		},
		settings: {
			badge: "Preferences",
			headline: <>Luminary,<br />your way.</>,
			sub: "Control notifications, security, and privacy all in one place.",
		},
	};
	const c = content[activeTab];

	return (
		<div className="hidden lg:flex lg:w-1/2 relative overflow-hidden"
			style={{ background: panelGradient }}>
			<div className="absolute -top-16 -right-16 w-80 h-80 rounded-full" style={{ background: "rgba(255,255,255,0.07)" }} />
			<div className="absolute -bottom-20 -left-20 w-96 h-96 rounded-full" style={{ background: "rgba(255,150,50,0.12)" }} />
			<div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-64 h-64 rounded-full" style={{ background: "rgba(255,255,255,0.05)" }} />

			<div className="relative z-10 flex flex-col justify-center px-12 py-14">
				<div className="inline-flex items-center gap-2 self-start rounded-full px-4 py-1.5 mb-8 text-xs font-medium text-white/90 border border-white/20"
					style={{ background: "rgba(255,255,255,0.12)", backdropFilter: "blur(4px)", fontFamily: poppins }}>
					<span className="w-1.5 h-1.5 rounded-full bg-green-400" />
					{c.badge}
				</div>

				<h2 className="text-2xl font-bold text-white leading-snug tracking-tight mb-3"
					style={{ fontFamily: poppins }}>{c.headline}</h2>
				<p className="text-sm text-white/70 leading-relaxed mb-8" style={{ fontFamily: poppins }}>{c.sub}</p>

				{/* SVG illustration */}
				<svg viewBox="0 0 340 200" fill="none" className="w-full max-w-sm">
					{/* Base card */}
					<rect x="20" y="30" width="300" height="140" rx="20" fill="rgba(255,255,255,0.08)" stroke="rgba(255,255,255,0.18)" strokeWidth="1" />
					{/* Avatar circle */}
					<circle cx="80" cy="90" r="32" fill="rgba(124,58,237,0.5)" stroke="rgba(255,255,255,0.3)" strokeWidth="1.5" />
					<circle cx="80" cy="82" r="12" fill="rgba(255,255,255,0.25)" />
					<ellipse cx="80" cy="106" rx="18" ry="10" fill="rgba(255,255,255,0.18)" />
					{/* Lines */}
					<rect x="128" y="70" width="100" height="8" rx="4" fill="rgba(255,255,255,0.3)" />
					<rect x="128" y="86" width="70" height="6" rx="3" fill="rgba(255,255,255,0.15)" />
					<rect x="128" y="100" width="85" height="6" rx="3" fill="rgba(255,255,255,0.15)" />
					{/* Edit badge */}
					<circle cx="284" cy="48" r="18" fill="rgba(249,115,22,0.6)" stroke="rgba(255,255,255,0.3)" strokeWidth="1" />
					<path d="M278 48l3 3 6-6" stroke="white" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" />
					{/* Bar indicators */}
					<rect x="40" y="148" width="260" height="6" rx="3" fill="rgba(255,255,255,0.08)" />
					<rect x="40" y="148" width="180" height="6" rx="3" fill="rgba(192,38,211,0.6)" />
					{/* Shield badge */}
					<circle cx="56" cy="48" r="14" fill="rgba(74,222,128,0.4)" stroke="rgba(255,255,255,0.3)" strokeWidth="1" />
					<path d="M56 42l5 2.5v4c0 2.5-2 4.5-5 5-3-0.5-5-2.5-5-5v-4L56 42Z" fill="rgba(255,255,255,0.7)" />
				</svg>

				{/* Stat cards */}
				<div className="flex gap-3 mt-7">
					{[{ num: "40K+", label: "Active users" }, { num: "99.9%", label: "Uptime SLA" }, { num: "4.9★", label: "Rating" }]
						.map(({ num, label }) => (
							<div key={label} className="flex-1 rounded-2xl px-4 py-3 border border-white/20"
								style={{ background: "rgba(255,255,255,0.1)", backdropFilter: "blur(6px)" }}>
								<div className="text-xl font-bold text-white tracking-tight" style={{ fontFamily: poppins }}>{num}</div>
								<div className="text-xs text-white/60 mt-0.5" style={{ fontFamily: poppins }}>{label}</div>
							</div>
						))}
				</div>
			</div>
		</div>
	);
}

// ══════════════════════════════════════════════════════════════════════════════
// TAB PANELS
// ══════════════════════════════════════════════════════════════════════════════

function ProfileTab({ toast }) {
	const [fullName, setFullName] = useState('');
	const [memberSince, setMemberSince] = useState('');

	const [form, setForm] = useState({ first_name: '', last_name: '', email: '', phone: '' });
	const set = k => e => setForm(f => ({ ...f, [k]: e.target.value }));

	const [loading, setLoading] = useState(false);
	const [modal, setModal] = useState({message: '', type: null});

	const navigate = useNavigate();

	const getUserDetails = async () => {
		try {
			setLoading(true);

			const result = await apiRequest(`${BASE_URL}/priv/user`, {auth: true});
			if(result) {
				setFullName(`${result.first_name} ${result.last_name}`);
				setMemberSince(formatDate(new Date(result.created_at)));
				setForm({
					first_name: result.first_name,
					last_name: result.last_name,
					email: result.email,
					phone: result.phone,
				});
			}
		} catch(err) {
			navigate('/signin');
		}

		setLoading(false);
	}

	const updateUserDetails = async () => {
		try {
			setLoading(true);

			const formData = {
				first_name: String(form.first_name).trim(),
				last_name: String(form.last_name).trim(),
				email: String(form.email).trim(),
				phone: String(form.phone).trim(),
			}

			for(const [field, value] of Object.entries(formData)) {
				if(!value) {
					const fieldName = `${field[0].toUpperCase()}${field.slice(1).toLowerCase()}`.replaceAll('_', ' ');
					throw new Error(`${fieldName} is empty`);
				}
			}

			const result = await apiRequest(
				`${BASE_URL}/priv/me`,
				{
					method: 'PATCH',
					body: JSON.stringify(formData),
					auth: true,
				}
			);
			if(result) {
				setFullName(`${result.first_name} ${result.last_name}`);
				setMemberSince(formatDate(new Date(result.created_at)));
				setForm({
					first_name: result.first_name,
					last_name: result.last_name,
					email: result.email,
					phone: result.phone,
				});

				toast("Profile updated!");
			}
		} catch(err) {
			const errorMessage = err?.message || 'Something went wrong';
			console.log("Failed to update info: ", errorMessage);
			setModal({message: errorMessage, type: 'error'});
			setTimeout(() => {
				setModal({message: '', type: null});
			}, 1000);
		} finally {
			setLoading(false);
		}
	}

	useEffect(() => {
		getUserDetails();
	}, []);

	return (
		<div className="space-y-6">
			{/* Avatar */}
			<div className="flex items-center gap-5 mb-2">
				<div className="relative">
					<div className="w-16 h-16 rounded-2xl flex items-center justify-center text-white text-xl font-bold"
						style={{ background: gradient }}>AR</div>
					<button className="absolute -bottom-1 -right-1 w-6 h-6 rounded-lg bg-white shadow flex items-center justify-center border border-gray-100">
						<Camera size={12} style={{ color: "#7C3AED" }} />
					</button>
				</div>
				<div>
					<p className="text-sm font-semibold text-gray-900" style={{ fontFamily: poppins }}>
						{fullName}
					</p>
					<p className="text-xs text-gray-400" style={{ fontFamily: poppins }}>
						{memberSince}
					</p>
				</div>
			</div>

			<div className="grid grid-cols-2 gap-5">
				<PrimaryField label="First Name" value={form.first_name} onChange={set('first_name')} />
				<PrimaryField label="Last Name" value={form.last_name} onChange={set("last_name")} />
			</div>
			<PrimaryField label="Email" type="email" value={form.email} onChange={set("email")} />
			<PrimaryField label="Phone" type="tel" value={form.phone} onChange={set("phone")} />
			{/*<PrimaryField label="Bio" value={form.bio} onChange={set("bio")} placeholder="A little about you…" />*/}

			<GradientButton className="w-full" onClick={updateUserDetails}>
				Save Changes →
			</GradientButton>

			{loading && <Loading />}

			{(modal?.message && modal?.type === 'error') && (
				<ErrorModal
					header="Oops!"
					message={modal.message}
					callback={() => {
						setModal({message: '', type: null});
					}}
				/>
			)}
		</div>
	);
}

function AddressTab({ toast }) {
	const [addresses, setAddresses] = useState([
		{ id: 1, label: "Home", line1: "123 Maple St", city: "San Francisco", state: "CA", zip: "94110", default: true },
		{ id: 2, label: "Work", line1: "456 Market Ave", city: "San Francisco", state: "CA", zip: "94105", default: false },
	]);
	const [adding, setAdding] = useState(false);
	const [form, setForm] = useState({ label: "", line1: "", city: "", state: "", zip: "" });
	const set = k => e => setForm(f => ({ ...f, [k]: e.target.value }));

	const addAddress = () => {
		if (!form.line1) return;
		setAddresses(a => [...a, { ...form, id: Date.now(), default: false }]);
		setForm({ label: "", line1: "", city: "", state: "", zip: "" });
		setAdding(false);
		toast("Address saved!");
	};

	const remove = id => setAddresses(a => a.filter(x => x.id !== id));
	const setDefault = id => setAddresses(a => a.map(x => ({ ...x, default: x.id === id })));

	return (
		<div className="space-y-4">
			{addresses.map(a => (
				<div key={a.id} className={`rounded-2xl border p-4 ${a.default ? "border-purple-200 bg-purple-50/40" : "border-gray-100 bg-gray-50/40"}`}>
					<div className="flex items-start justify-between">
						<div>
							<div className="flex items-center gap-2 mb-1">
								<span className="text-xs font-semibold uppercase tracking-widest text-gray-400" style={{ fontFamily: poppins }}>{a.label || "Address"}</span>
								{a.default && <span className="text-[10px] px-2 py-0.5 rounded-full text-purple-600 bg-purple-100 font-semibold" style={{ fontFamily: poppins }}>Default</span>}
							</div>
							<p className="text-sm text-gray-800" style={{ fontFamily: poppins }}>{a.line1}</p>
							<p className="text-xs text-gray-400" style={{ fontFamily: poppins }}>{a.city}, {a.state} {a.zip}</p>
						</div>
						<div className="flex gap-2">
							{!a.default && (
								<button onClick={() => setDefault(a.id)} className="text-xs text-purple-600 font-semibold hover:text-purple-700 transition-colors" style={{ fontFamily: poppins }}>
									Set default
								</button>
							)}
							<button onClick={() => remove(a.id)} className="text-gray-300 hover:text-red-400 transition-colors">
								<Trash2 size={14} />
							</button>
						</div>
					</div>
				</div>
			))}

			{adding ? (
				<div className="rounded-2xl border border-gray-100 p-5 space-y-4 bg-gray-50/40">
					<div className="grid grid-cols-2 gap-4">
						<PrimaryField label="Label" value={form.label} onChange={set("label")} placeholder="Home / Work…" />
						<PrimaryField label="Street" value={form.line1} onChange={set("line1")} placeholder="123 Main St" />
					</div>
					<div className="grid grid-cols-3 gap-4">
						<PrimaryField label="City" value={form.city} onChange={set("city")} placeholder="City" />
						<PrimaryField label="State" value={form.state} onChange={set("state")} placeholder="CA" />
						<PrimaryField label="ZIP" value={form.zip} onChange={set("zip")} placeholder="94110" />
					</div>
					<div className="flex gap-3">
						<GradientButton className="flex-1" onClick={addAddress}>Save Address →</GradientButton>
						<button onClick={() => setAdding(false)} className="text-xs text-gray-400 hover:text-gray-600 transition-colors" style={{ fontFamily: poppins }}>Cancel</button>
					</div>
				</div>
			) : (
				<button onClick={() => setAdding(true)}
					className="w-full py-3 rounded-2xl border-2 border-dashed border-gray-200 text-xs font-semibold text-gray-400 hover:border-purple-300 hover:text-purple-500 transition-all flex items-center justify-center gap-2"
					style={{ fontFamily: poppins }}>
					<Plus size={14} /> Add New Address
				</button>
			)}
		</div>
	);
}

function PaymentTab({ toast }) {
	const [cards, setCards] = useState([
		{ id: 1, brand: "Visa", last4: "4242", expiry: "12/26", default: true },
		{ id: 2, brand: "Mastercard", last4: "8888", expiry: "09/25", default: false },
	]);
	const [adding, setAdding] = useState(false);
	const [form, setForm] = useState({ number: "", name: "", expiry: "", cvc: "" });
	const set = k => e => setForm(f => ({ ...f, [k]: e.target.value }));

	const CardIcon = ({ brand }) => (
		<div className={`w-8 h-5 rounded flex items-center justify-center text-[9px] font-bold text-white ${brand === "Visa" ? "bg-blue-600" : "bg-orange-500"}`}
			style={{ fontFamily: poppins }}>{brand === "Visa" ? "VISA" : "MC"}</div>
	);

	const remove = id => setCards(c => c.filter(x => x.id !== id));
	const setDefault = id => setCards(c => c.map(x => ({ ...x, default: x.id === id })));

	return (
		<div className="space-y-4">
			{cards.map(c => (
				<div key={c.id} className={`rounded-2xl border p-4 ${c.default ? "border-purple-200 bg-purple-50/40" : "border-gray-100 bg-gray-50/40"}`}>
					<div className="flex items-center justify-between">
						<div className="flex items-center gap-3">
							<CardIcon brand={c.brand} />
							<div>
								<p className="text-sm font-semibold text-gray-900" style={{ fontFamily: poppins }}>•••• •••• •••• {c.last4}</p>
								<p className="text-xs text-gray-400" style={{ fontFamily: poppins }}>Expires {c.expiry}</p>
							</div>
						</div>
						<div className="flex items-center gap-3">
							{c.default && <span className="text-[10px] px-2 py-0.5 rounded-full text-purple-600 bg-purple-100 font-semibold" style={{ fontFamily: poppins }}>Default</span>}
							{!c.default && <button onClick={() => setDefault(c.id)} className="text-xs text-purple-600 font-semibold" style={{ fontFamily: poppins }}>Set default</button>}
							<button onClick={() => remove(c.id)} className="text-gray-300 hover:text-red-400 transition-colors"><Trash2 size={14} /></button>
						</div>
					</div>
				</div>
			))}

			{adding ? (
				<div className="rounded-2xl border border-gray-100 p-5 space-y-4 bg-gray-50/40">
					<PrimaryField label="Card Number" value={form.number} onChange={set("number")} placeholder="1234 5678 9012 3456" />
					<PrimaryField label="Cardholder Name" value={form.name} onChange={set("name")} placeholder="Alex Rivera" />
					<div className="grid grid-cols-2 gap-4">
						<PrimaryField label="Expiry" value={form.expiry} onChange={set("expiry")} placeholder="MM/YY" />
						<PrimaryField label="CVC" value={form.cvc} onChange={set("cvc")} placeholder="•••" />
					</div>
					<div className="flex items-center gap-2 text-xs text-gray-400 mb-2" style={{ fontFamily: poppins }}>
						<Shield size={12} className="text-green-500" /> Your card details are encrypted & secure.
					</div>
					<div className="flex gap-3">
						<GradientButton className="flex-1" onClick={() => { setAdding(false); toast("Card added!"); }}>Add Card →</GradientButton>
						<button onClick={() => setAdding(false)} className="text-xs text-gray-400 hover:text-gray-600 transition-colors" style={{ fontFamily: poppins }}>Cancel</button>
					</div>
				</div>
			) : (
				<button onClick={() => setAdding(true)}
					className="w-full py-3 rounded-2xl border-2 border-dashed border-gray-200 text-xs font-semibold text-gray-400 hover:border-purple-300 hover:text-purple-500 transition-all flex items-center justify-center gap-2"
					style={{ fontFamily: poppins }}>
					<Plus size={14} /> Add New Card
				</button>
			)}
		</div>
	);
}

function OrdersTab() {
	const [expanded, setExpanded] = useState(null);

	return (
		<div className="space-y-3">
			{mockOrders.map(o => (
				<div key={o.id} className="rounded-2xl border border-gray-100 overflow-hidden">
					<button
						onClick={() => setExpanded(expanded === o.id ? null : o.id)}
						className="w-full px-4 py-4 flex items-center justify-between hover:bg-gray-50/60 transition-colors"
					>
						<div className="flex items-center gap-3">
							<div className="w-9 h-9 rounded-xl flex items-center justify-center bg-gray-100">
								<Package size={16} className="text-gray-500" />
							</div>
							<div className="text-left">
								<p className="text-sm font-semibold text-gray-900" style={{ fontFamily: poppins }}>{o.id}</p>
								<p className="text-xs text-gray-400" style={{ fontFamily: poppins }}>{o.date} · {o.items} item{o.items > 1 ? "s" : ""}</p>
							</div>
						</div>
						<div className="flex items-center gap-3">
							<span className={`text-[10px] font-semibold px-2.5 py-1 rounded-full ${statusColor[o.status]}`} style={{ fontFamily: poppins }}>{o.status}</span>
							<span className="text-sm font-semibold text-gray-900" style={{ fontFamily: poppins }}>{o.total}</span>
							<ChevronDown size={14} className={`text-gray-400 transition-transform ${expanded === o.id ? "rotate-180" : ""}`} />
						</div>
					</button>

					{expanded === o.id && (
						<div className="border-t border-gray-100 px-4 py-4 bg-gray-50/40 space-y-3">
							<div className="flex items-center gap-2 text-xs text-gray-500" style={{ fontFamily: poppins }}>
								<div className="w-8 h-8 rounded-xl bg-gray-200 flex items-center justify-center text-gray-400 text-[10px]">IMG</div>
								<span>Product details would appear here — connect to your order model.</span>
							</div>
							<div className="flex items-center justify-between">
								<div>
									{o.rating ? (
										<div className="flex items-center gap-1">
											{Array.from({ length: 5 }).map((_, i) => (
												<Star key={i} size={12} fill={i < o.rating ? "#f97316" : "none"} stroke={i < o.rating ? "#f97316" : "#cbd5e1"} />
											))}
											<span className="text-xs text-gray-400 ml-1" style={{ fontFamily: poppins }}>Your rating</span>
										</div>
									) : (
										<span className="text-xs text-gray-400" style={{ fontFamily: poppins }}>Not yet rated</span>
									)}
								</div>
								<button className="text-xs text-purple-600 font-semibold hover:text-purple-700 transition-colors flex items-center gap-1" style={{ fontFamily: poppins }}>
									Reorder <ChevronRight size={12} />
								</button>
							</div>
						</div>
					)}
				</div>
			))}
		</div>
	);
}

function SettingsTab({ toast }) {
	const [notifications, setNotifications] = useState({ orders: true, promos: false, security: true });
	const [privacy, setPrivacy] = useState("friends");

	const Toggle = ({ checked, onChange }) => (
		<button onClick={() => onChange(!checked)}
			className={`relative w-10 h-5 rounded-full transition-all duration-200 ${checked ? "" : "bg-gray-200"}`}
			style={checked ? { background: gradient } : {}}>
			<span className={`absolute top-0.5 w-4 h-4 rounded-full bg-white shadow transition-all duration-200 ${checked ? "left-5" : "left-0.5"}`} />
		</button>
	);

	const Row = ({ icon: Icon, label, children }) => (
		<div className="flex items-center justify-between py-3 border-b border-gray-100 last:border-0">
			<div className="flex items-center gap-3">
				<div className="w-7 h-7 rounded-lg bg-gray-100 flex items-center justify-center">
					<Icon size={13} className="text-gray-500" />
				</div>
				<span className="text-sm text-gray-700" style={{ fontFamily: poppins }}>{label}</span>
			</div>
			{children}
		</div>
	);

	return (
		<div className="space-y-6">
			{/* Notifications */}
			<div>
				<p className="text-xs font-semibold tracking-widest text-gray-400 uppercase mb-3" style={{ fontFamily: poppins }}>Notifications</p>
				<div className="rounded-2xl border border-gray-100 px-4">
					<Row icon={ShoppingBag} label="Order updates">
						<Toggle checked={notifications.orders} onChange={v => setNotifications(n => ({ ...n, orders: v }))} />
					</Row>
					<Row icon={Bell} label="Promotions & offers">
						<Toggle checked={notifications.promos} onChange={v => setNotifications(n => ({ ...n, promos: v }))} />
					</Row>
					<Row icon={Shield} label="Security alerts">
						<Toggle checked={notifications.security} onChange={v => setNotifications(n => ({ ...n, security: v }))} />
					</Row>
				</div>
			</div>

			{/* Security */}
			<div>
				<p className="text-xs font-semibold tracking-widest text-gray-400 uppercase mb-3" style={{ fontFamily: poppins }}>Security</p>
				<div className="rounded-2xl border border-gray-100 px-4">
					<Row icon={Lock} label="Change password">
						<button className="text-xs text-purple-600 font-semibold hover:text-purple-700 transition-colors flex items-center gap-1" style={{ fontFamily: poppins }}>
							Update <ChevronRight size={12} />
						</button>
					</Row>
					<Row icon={Shield} label="Two-factor auth">
						<button className="text-xs text-purple-600 font-semibold hover:text-purple-700 transition-colors flex items-center gap-1" style={{ fontFamily: poppins }}>
							Enable <ChevronRight size={12} />
						</button>
					</Row>
				</div>
			</div>

			{/* Danger zone */}
			<div>
				<p className="text-xs font-semibold tracking-widest text-gray-400 uppercase mb-3" style={{ fontFamily: poppins }}>Account</p>
				<div className="rounded-2xl border border-gray-100 px-4">
					<Row icon={LogOut} label="Sign out">
						<button onClick={() => toast("Signed out!")} className="text-xs text-red-400 font-semibold hover:text-red-500 transition-colors flex items-center gap-1" style={{ fontFamily: poppins }}>
							Sign out <ChevronRight size={12} />
						</button>
					</Row>
				</div>
			</div>
		</div>
	);
}

// ══════════════════════════════════════════════════════════════════════════════
// MAIN PAGE
// ══════════════════════════════════════════════════════════════════════════════

const Profile = () => {
	const [active, setActive] = useState("profile");
	const [toastMsg, setToastMsg] = useState("");

	const toast = msg => {
		setToastMsg(msg);
		setTimeout(() => setToastMsg(""), 2800);
	};

	const renderTab = () => {
		switch (active) {
			case "profile": return <ProfileTab toast={toast} />;
			case "address": return <AddressTab toast={toast} />;
			case "payment": return <PaymentTab toast={toast} />;
			case "orders": return <OrdersTab />;
			case "settings": return <SettingsTab toast={toast} />;
		}
	};

	return (
		<div className="bg-gray-100 min-h-screen flex items-center justify-center p-4" style={{ fontFamily: poppins }}>
			<div className="w-full h-full min-h-[calc(100vh-2rem)] rounded-3xl overflow-hidden shadow-2xl bg-white flex">

				{/* LEFT PANEL */}
				<div className="w-full px-8 sm:px-10 py-12 flex flex-col">
					<Logo />

					{/* Tab nav */}
					<div className="flex gap-1 mb-8 bg-gray-50 rounded-2xl p-1">
						{TABS.map(({ id, label, icon: Icon }) => (
							<button
								key={id}
								onClick={() => setActive(id)}
								className={`flex-1 flex flex-col items-center gap-1 py-2.5 rounded-xl text-[10px] font-semibold tracking-wide transition-all duration-200 ${active === id ? "bg-white shadow text-purple-600" : "text-gray-400 hover:text-gray-600"}`}
								style={{ fontFamily: poppins }}
							>
								<Icon size={15} />
								<span className="hidden sm:block">{label}</span>
							</button>
						))}
					</div>

					{/* Tab heading */}
					<div className="mb-6">
						<h1 className="text-2xl font-bold text-gray-900 leading-tight tracking-tight" style={{ fontFamily: poppins }}>
							{TABS.find(t => t.id === active)?.label}
						</h1>
						<p className="text-xs text-gray-400 mt-1" style={{ fontFamily: poppins }}>
							Manage your {TABS.find(t => t.id === active)?.label.toLowerCase()} details
						</p>
					</div>

					{/* Tab content */}
					<div className="flex-1 overflow-y-auto pr-1">
						{renderTab()}
					</div>
				</div>

				{/* RIGHT PANEL */}
				{/*<RightPanel activeTab={active} />*/}
			</div>

			<Toast msg={toastMsg} />
		</div>
	);
}

export default Profile;