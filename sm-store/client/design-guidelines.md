# Luminary — Design System Prompt

## Master Prompt Template

> Build a **[PAGE NAME]** page for a ecommerce web app called **Luminary** using **React + Tailwind CSS**.
> Follow the Luminary design system described below exactly.

---

## 1. Layout

- **Split-screen layout** — two columns side by side inside a `max-w-5xl` centered container for sign in, sign up, forgot password page and other minimal content pages.
- Outer wrapper: `rounded-3xl overflow-hidden shadow-2xl bg-white`, `min-height: calc(100vh-1rem) and min-width: calc(100vw-1rem) for pages with dense contents like dashboards or user profile page. Don't apply this outer wrapper for home page and other pages with text content`.
- Page background for outer wrapper: light gray (`bg-gray-100`), centered with `min-h-screen flex items-center justify-center p-4 and padding of 1rem to give space for outer wrapper`.
- **Left panel** (`lg:w-1/2`): white background, `px-10 py-12`, flex column, vertically centered content. Contains the form and other minimal content.
- **Right panel** (`hidden lg:flex lg:w-1/2`): gradient background, decorative blobs, SVG illustration, stat cards. Always visible only on `lg` and above; hidden on mobile. This is also for minimal content that might contains form.

---

## 2. Typography

- **Font**: Poppins (Google Fonts) — import {poppins} from '@/config/style'
- Apply Poppins inline on the root wrapper and all inputs/buttons.
- **Headings** (`h1`): `text-3xl font-bold text-gray-900 leading-tight tracking-tight`
- **Subheading / tagline**: `text-sm text-gray-400`
- **Field labels**: `text-xs font-semibold tracking-widest text-gray-400 uppercase`
- **Body / helper text**: `text-xs text-gray-400`
- **Links**: `text-purple-600 font-semibold hover:text-purple-700 transition-colors`

---

## 3. Color Palette

| Role | Value |
|---|---|
| Primary purple | `#7C3AED` |
| Pink / magenta | `#C026D3` |
| Orange accent | `#f97316` |
| Dark navy (text) | `#1a1a2e` |
| Muted label | `#8892a4` / `text-gray-400` |
| Border / divider | `#e2e8f0` / `border-gray-200` |
| Input placeholder | `#cbd5e1` / `placeholder-gray-300` |
| Success green | `#22c55e` |
| Error red | `#ef4444` |
| Warning orange | `#f97316` |

---

## 4. Gradient

The brand gradient is used on the background:

```
background: linear-gradient(160deg, #4f1899 0%, #7C3AED 30%, #C026D3 62%, #f97316 100%);
```

---

## 5. Components

### Logo (top-left of left panel)
```jsx
<div className="flex items-center gap-2 mb-10">
  <div className="w-8 h-8 rounded-xl flex items-center justify-center"
    style={{ background: "linear-gradient(135deg, #7C3AED, #C026D3)" }}>
    {/* 14x14 white SVG icon */}
  </div>
  <span className="text-base font-semibold text-gray-900 tracking-tight">Luminary</span>
</div>
```

### Form Fields (underline style)

```jsx
export function PrimaryField({ label, id, type="text", value, onChange, placeholder, disabled, required=false })
```

Import and use it like this.
```jsx
<PrimaryField label="Name" value={form.name} onChange={set('name')} />
```

### CTA Button (gradient, rounded)
```jsx
export function GradientButton({ children, onClick, disabled, type = "button", className = "" })
```

import it and use it like this.
```jsx
<GradientButton className="w-full" onClick={updateUserDetails}>
  Save Changes →
</GradientButton>
```

### OR Divider
```jsx
<div className="flex items-center gap-3 my-5">
  <div className="flex-1 h-px bg-gray-100" />
  <span className="text-xs text-gray-300">or continue with</span>
  <div className="flex-1 h-px bg-gray-100" />
</div>
```

### Bottom Helper Text
```jsx
<p className="text-center text-xs text-gray-400 mt-6">
  Helper text?{" "}
  <a href="#" className="text-purple-600 font-semibold hover:text-purple-700 transition-colors">
    Action link →
  </a>
</p>
```

---

## 6. Right Panel Structure (Split Screen Layout)

```jsx
<div className="hidden lg:flex lg:w-1/2 relative overflow-hidden"
  style={{ background: "linear-gradient(160deg, #4f1899 0%, #7C3AED 30%, #C026D3 62%, #f97316 100%)" }}>

  {/* Decorative blobs */}
  <div className="absolute -top-16 -right-16 w-80 h-80 rounded-full"
    style={{ background: "rgba(255,255,255,0.07)" }} />
  <div className="absolute -bottom-20 -left-20 w-96 h-96 rounded-full"
    style={{ background: "rgba(255,150,50,0.12)" }} />
  <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-64 h-64 rounded-full"
    style={{ background: "rgba(255,255,255,0.05)" }} />

  {/* Content */}
  <div className="relative z-10 flex flex-col justify-center px-12 py-14">

    {/* Status badge */}
    <div className="inline-flex items-center gap-2 self-start rounded-full px-4 py-1.5 mb-8
      text-xs font-medium text-white/90 border border-white/20"
      style={{ background: "rgba(255,255,255,0.12)", backdropFilter: "blur(4px)" }}>
      <span className="w-1.5 h-1.5 rounded-full bg-green-400" />
      Badge text here
    </div>

    {/* Headline */}
    <h2 className="text-2xl font-bold text-white leading-snug tracking-tight mb-3">
      Headline text<br />goes here.
    </h2>
    <p className="text-sm text-white/70 leading-relaxed mb-8">
      Supporting description copy goes here.
    </p>

    {/* SVG Illustration — custom per page, follows rules below */}
    <IllustrationSVG />

    {/* Stat cards */}
    <div className="flex gap-3 mt-7">
      {[
        { num: "40K+", label: "Active teams" },
        { num: "99.9%", label: "Uptime SLA" },
        { num: "4.9★", label: "User rating" },
      ].map(({ num, label }) => (
        <div key={label} className="flex-1 rounded-2xl px-4 py-3 border border-white/20"
          style={{ background: "rgba(255,255,255,0.1)", backdropFilter: "blur(6px)" }}>
          <div className="text-xl font-bold text-white tracking-tight">{num}</div>
          <div className="text-xs text-white/60 mt-0.5">{label}</div>
        </div>
      ))}
    </div>
  </div>
</div>
```

---

## 7. SVG Illustration Rules (Right Panel)

Minimal page gets a unique contextual SVG illustration. All illustrations follow these rules:

- `viewBox="0 0 340 200"` (or `210` if taller), `fill="none"`, `className="w-full max-w-sm"`
- All shapes use `rgba(255,255,255,...)` or `rgba(r,g,b,0.x)` fills — semi-transparent white or brand colors
- Stroke: `rgba(255,255,255,0.18)` to `rgba(255,255,255,0.3)`, `strokeWidth="1"` or `1.5`
- Floating badges: small circles (`r="10"–"24"`) with icons inside, positioned at corners
- Main content: a frosted card (`rgba(255,255,255,0.08)`) with context-relevant elements inside
- Accent colors allowed: orange `rgba(249,115,22,0.5)`, purple `rgba(124,58,237,0.6)`, pink `rgba(192,38,211,0.6)`, blue `rgba(96,165,250,0.7)`, green `rgba(74,222,128,0.8)`
- No text inside SVG illustrations (use shapes only)

**Page-specific illustration ideas:**
| Page | Illustration Concept |
|---|---|
| Login | UI dashboard card with floating check/add badges |
| Signup | Team avatars row, progress bar, task checklist |
| Forgot Password | Envelope icon, key shape, lock circle |
| Pricing | Tier cards, checkmarks, crown badge |
| 404 / Error | Broken path, warning triangle, refresh badge |

---

## 8. Spacing & Radius Reference

| Token | Value |
|---|---|
| Card border radius | `rounded-3xl` (outer), `rounded-2xl` (buttons, stat cards) |
| Input fields | underline only — `border-b border-gray-200` |
| Button padding | `py-3.5` full-width |
| Panel padding | `px-10 py-12` (left), `px-12 py-14` (right) |
| Field spacing | `space-y-5` or `space-y-6` |
| Logo margin bottom | `mb-10` or `mb-12` |
| Heading margin bottom | `mb-1` (h1) + `mb-8` (tagline to form) |

---

## 9. Responsive Behavior

- Mobile: only the left (form) panel is shown. Right panel uses `hidden lg:flex`.
- Left panel fills full width on mobile: `w-full lg:w-1/2`
- Reduce padding on mobile if needed: `px-6 sm:px-10`

---

## 10. Page Checklist

When building a new page, make sure it includes:

- [ ] Luminary logo (top-left, gradient mark + wordmark)
- [ ] Bold h1 heading
- [ ] Muted subtitle / tagline below heading
- [ ] Form fields with uppercase labels + underline inputs + purple focus state
- [ ] Gradient CTA button with hover lift + active scale
- [ ] Helper/navigation link at the bottom (`text-xs text-gray-400`)
- [ ] Right panel: gradient bg + blobs + badge + headline + illustration + 3 stat cards
- [ ] Poppins font applied globally via inline style on root wrapper
- [ ] All inputs and buttons have `fontFamily: "'Poppins', sans-serif"` set explicitly

---

## Quick Update

- Use lucide-react icons
- Only applies split screen layout on pages with form and any page that needs user to focus at the center for action like a modal.
- Add blobs for full width layout pages like dashboard, profile and product display with dense content. I need the blob to be semi-transparent with the touch of gradient style above.


## Database tables and columns

-- product

product_category
  - parent_category_id
  - category_name
  - id

product
  - category_id
  - name
  - description
  - product_image
  - id

variation
  - category_id
  - name
  - id

variation_option
  - variation_id
  - value
  - id

product_item
  - product_id
  - sku
  - qty_in_stock
  - product_image
  - price
  - created_at
  - updated_at
  - id

product_configuration
  - product_item_id
  - variation_option_id

-- payment

payment_type
  - value
  - id

user_payment_method
  - user_id
  - payment_type_id
  - provider
  - credit_card_number
  - credit_card_exp_month
  - credit_card_exp_day
  - credit_card_secret_code
  - name_on_card
  - is_default
  - created_at
  - updated_at
  - id

-- shopping cart

shopping_cart
  - user_id
  - id

shopping_cart_item
  - cart_id
  - product_item_id
  - qty
  - created_at
  - updated_at
  - id

-- order

shipping_method
  - name
  - price
  - id

order_status
  - status
  - created_at
  - id

shop_order
  - user_id
  - order_date
  - purchase_order_number
  - payment_method_id
  - shipping_address
  - shipping_method
  - order_total
  - order_status
  - id

-- order line

order_line
  - product_item_id
  - order_id
  - qty
  - price
  - id

-- reviews

user_review
  - user_id
  - ordered_product_id
  - rating_value
  - comment
  - created_at
  - updated_at
  - id

-- promotion

promotion
  - name
  - description
  - discount_rate
  - start_date
  - end_date
  - id

promotion_category
  - category_id
  - promotion_id