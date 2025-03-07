<div dir="rtl">

# HTTP چیست؟

> HTTP یک پروتکل قابل گسترش است که در طول زمان تکامل یافته است. این یک پروتکل لایه‌ی کاربرد است که از طریق TCP یا یک اتصال TCP رمزگذاری‌شده با TLS ارسال می‌شود، هرچند که از نظر تئوری می‌توان از هر پروتکل انتقال قابل‌اعتمادی استفاده کرد. به دلیل قابلیت گسترش آن، HTTP نه‌تنها برای دریافت اسناد ابرمتنی (Hypertext) بلکه برای دریافت تصاویر و ویدیوها یا ارسال محتوا به سرورها، مانند نتایج فرم‌های HTML، مورد استفاده قرار می‌گیرد. همچنین می‌توان از HTTP برای دریافت بخش‌هایی از اسناد جهت به‌روزرسانی صفحات وب به صورت پویا استفاده کرد.

## اجزای سیستم‌های مبتنی بر HTTP

### کلاینت (Client)

> عامل کاربر (User-Agent) هر ابزاری است که از طرف کاربر عمل می‌کند. این نقش عمدتاً توسط مرورگر وب انجام می‌شود.
>
> برای نمایش یک صفحه‌ی وب، مرورگر یک درخواست اولیه برای دریافت سند HTML که نمایانگر صفحه است ارسال می‌کند. سپس این فایل را پردازش کرده و درخواست‌های اضافی برای اجرای اسکریپت‌ها، اطلاعات مربوط به چیدمان (CSS) و منابع جانبی درون صفحه (معمولاً تصاویر و ویدیوها) را ارسال می‌کند.

### سرور وب (Web Server)

> در طرف مقابل کانال ارتباطی، سروری قرار دارد که اسناد درخواستی را به کلاینت ارائه می‌دهد.
> سرور از نظر ظاهری به عنوان یک ماشین منفرد دیده می‌شود؛ اما در واقع می‌تواند مجموعه‌ای از سرورها باشد که بار را بین یکدیگر تقسیم می‌کنند (Load Balancing)، یا شامل نرم‌افزارهای دیگری مانند کش، سرور پایگاه داده، یا سرورهای تجارت الکترونیک باشد که به‌طور کامل یا جزئی اسناد را بر اساس درخواست ایجاد می‌کنند.

### پروکسی‌ها بین کلاینت و سرور (Proxies)

> در واقعیت، بین مرورگر وب و سروری که درخواست را پردازش می‌کند، تعداد زیادی کامپیوتر و ماشین وجود دارند که پیام‌های HTTP را منتقل می‌کنند. به دلیل ساختار لایه‌ای وب، بیشتر این دستگاه‌ها در لایه‌های شبکه و انتقال عمل کرده و در سطح HTTP نامرئی هستند، اما می‌توانند تأثیر قابل‌توجهی بر عملکرد داشته باشند.
>
> دستگاه‌هایی که در سطح لایه‌ی کاربرد (Application Layer) عمل می‌کنند، معمولاً پروکسی نامیده می‌شوند. پروکسی‌ها می‌توانند به دو شکل باشند:
> - **شفاف (Transparent):** درخواست‌ها را بدون تغییر به سرور اصلی ارسال می‌کنند.
> - **غیرشفاف (Non-Transparent):** درخواست‌ها را پیش از ارسال به سرور تغییر می‌دهند.
>
> پروکسی‌ها می‌توانند چندین عملکرد مختلف داشته باشند، از جمله:
> - **کش (Caching):** کش می‌تواند عمومی یا خصوصی (مانند کش مرورگر) باشد.
> - **فیلترینگ (Filtering):** مانند اسکن ویروس یا کنترل والدین.
> - **بالانس بار (Load Balancing):** توزیع درخواست‌ها بین چندین سرور.
> - **احراز هویت (Authentication):** کنترل دسترسی به منابع مختلف.
> - **ثبت گزارشات (Logging):** ذخیره‌ی اطلاعات تاریخچه‌ای.

---

## HTTP بدون وضعیت (Stateless) در مقابل دارای نشست (Sessionful)

> HTTP به‌صورت ذاتی بدون وضعیت (Stateless) است، به این معنی که هیچ ارتباطی بین دو درخواست مختلف وجود ندارد. یعنی نمی‌توان تشخیص داد که دو درخواست از سوی یک کلاینت آمده‌اند یا خیر.
> اما HTTP می‌تواند حالت‌مند (Sessionful) شود؛ با استفاده از کوکی‌ها و برخی هدرها، می‌توان یک نشست مشترک بین درخواست‌های مختلف ایجاد کرد.

---

## اتصالات HTTP

> HTTP یک پروتکل لایه‌ی کاربرد است، به این معنی که نحوه‌ی انتقال آن را تعیین نمی‌کند.
> مدیریت اتصال و انتقال داده‌ها بر عهده‌ی لایه‌ی انتقال است.
> HTTP به یک پروتکل انتقال قابل‌اعتماد نیاز دارد که داده‌ها را در حین انتقال از بین نبرد.
>
> دو پروتکل رایج برای انتقال داده‌ها عبارتند از:
> - **TCP:** پروتکلی قابل‌اعتماد که تضمین می‌کند داده‌ها بدون خطا و به ترتیب صحیح دریافت شوند.
> - **UDP:** پروتلی که سریع‌تر است اما تضمین نمی‌کند که همه‌ی داده‌ها بدون از دست رفتن دریافت شوند.
>
> در ارتباطات HTTP، اتصال TCP می‌تواند به دو شکل باشد:
> - **غیرپایدار (Non-Persistent):** اتصال TCP برای هر درخواست باز شده و پس از دریافت پاسخ بسته می‌شود.
> - **پایدار (Persistent):** یک اتصال TCP ایجاد شده و برای چندین درخواست استفاده می‌شود تا زمانی که زمان انتظار آن به پایان برسد.

---

## جریان کاری HTTP

1. **باز کردن یک اتصال TCP**
2. **ارسال یک پیام HTTP**

```HTTP
GET / HTTP/1.1
Host: developer.mozilla.org
Accept-Language: fa
```

3. **دریافت پاسخ سرور**

```HTTP
HTTP/1.1 200 OK
Date: Sat, 09 Oct 2010 14:28:02 GMT
Server: Apache
Last-Modified: Tue, 01 Dec 2009 20:18:22 GMT
ETag: "51142bc1-7449-479b075b2891b"
Accept-Ranges: bytes
Content-Length: 29769
Content-Type: text/html

<!doctype html>… (اینجا محتوای صفحه وب قرار دارد)
```

4. **بستن یا استفاده مجدد از اتصال برای درخواست‌های بعدی.**

---

## پایپ‌لاینینگ (Pipelining)

> در صورت فعال بودن **پایپ‌لاینینگ HTTP**، می‌توان چندین درخواست را بدون انتظار برای دریافت پاسخ‌های قبلی ارسال کرد.
> HTTP/2 این مشکل را با معرفی **مولتی‌پلکسینگ (Multiplexing)** درون یک فریم برطرف کرده است.

---

## کدهای وضعیت (Status Codes)

> - **۱۰۰ - ۱۹۹:** پاسخ‌های اطلاعاتی
> - **۲۰۰ - ۲۹۹:** پاسخ‌های موفقیت‌آمیز
> - **۳۰۰ - ۳۹۹:** پاسخ‌های تغییر مسیر
> - **۴۰۰ - ۴۹۹:** خطاهای سمت کلاینت
> - **۵۰۰ - ۵۹۹:** خطاهای سمت سرور  

[لیست کامل کدهای وضعیت](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

</div>
