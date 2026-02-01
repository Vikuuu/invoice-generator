#let bill-to(
  name:"",
  gstin: "",
  address: "",
) = {
  pad(
    top: 0.25em,
    align(left)[
      *Bill To* \ \
      *#name* \ 
      *GSTIN*: #gstin \
      *Address*: #address \
    ]
  )
}

#let product-table(
  items: (),
) = {
  pad(
    top: 1.25em,
    table(
      columns: (1fr, 1fr, 1fr, 1fr, 1fr, 1fr),
      align: horizon,
      table.header(
        //..items.map(dictionary.keys).flatten()
        [*Item Name*], [*HSN/SAC*], [*Qty(m)*], [*Price*], [*GST*], [*Total*]
      ),
      ..items.map(dictionary.values).flatten(),
    )
  )
}

#let ship-to(
  address: "",
) = {
  pad(
    top: 0.25em,
    align(left)[
      *ship to*: #address
    ]
  )
}

#let payment-table(
  payment-data: (:),

  sub-total: decimal("0.00"),
  igst: decimal("0.00"),
  total: decimal("0.00"),
) = {
  pad(
    top: 0.25em,
    grid(
      columns: (1fr, 1fr),
      [
        *Payment to* \
        Account Holder: #payment-data.acc-name \
        Account Number: #payment-data.acc-number \
        IFSC: #payment-data.ifsc \
        Branch: #payment-data.branch \
        Bank Name: #payment-data.bank-name \ \ \
        Virtual Payment Address: #payment-data.virtual-address \ \ \
      ],
      [
        #grid(
          columns: (1fr, 1fr),
          align(left)[
            Sub Total \
            IGST(5%)  \ \ \
            Total \
          ],
          align(left)[
            ₹ #sub-total \
            ₹ #igst \ \ \
            ₹ #total \ 
          ],
        )
      ]
    )
  )
}

#let sign-image(
  image-path: "",
  wid: 100pt,
) = {
  let img = image(image-path, width: wid)
  show: pad(
    top: 0.25em,
    align(left)[
      == Thank you for your purchase!
      #img
      #line(length: wid)
      *Authorised Signator*
    ]
  )
}
#let invoice(
  company: "",
  company-pos: center,
  company-gstin: "",
  company-address: "",
  company-add-pos: center,
  invoice-date: datetime.today(), 
  invoice-number: int(0),
  bill-to-name: "",
  bill-to-gstin: "",
  bill-to-address: "",
  items: (),
  ship-to-address: "",
  payment-data: (:),
  sub-total: decimal("0.00"),
  igst: decimal("0.00"),
  image-path: "fake-sign.jpg",

  font: "New Computer Modern",
  company-font-size: 20pt,
  font-size: 10pt,
  lang: "en",
  paper: "a4",

  body,
) = {
  let filename = str("")
  filename = str(company) + datetime.today().display()
  set document(author: company, title: filename)

  set text(
    font: font,
    size: font-size,
    lang: lang,
    // Disable ligatures so ATS systems do not get
    // confused when parsing fonts.
    ligatures: false
  )

  set page(
    margin: (x: 2.8cm, y: 1.5cm),
    paper: paper,
  )

  show heading.where(level: 1): it => [
    #set align(company-pos)
    #set text(
      weight: 700,
      size: company-font-size,
    )
    #pad(it.body)
  ]

  [= #(company)]

  pad(
    top: 0.25em,
    align(company-add-pos)[
      *GSTIN: #(company-gstin)* \
      #(company-address)
    ],
  )

  bill-to(
    name: bill-to-name,
    gstin: bill-to-gstin,
    address: bill-to-address,
  )
  product-table(
    items: items,
  )
  ship-to(
    address: ship-to-address,
  )
  payment-table(
    payment-data: payment-data,
  )
  sign-image(
    image-path: image-path,
  )
}
