// see page 133

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+ call
	forEachNode(doc, startElement, endElement)
	//!- call

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// !+startend
var depth int

// startElement prints the opening HTML tag.
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

// endElement prints the closing HTML tag.
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend

// go build gopl.io/ch5/outline2
// ./outline2 http://gopl.io

// Output:
// <html>
//   <head>
//     <meta>
//     </meta>
//     <title>
//     </title>
//     <script>
//     </script>
//     <link>
//     </link>
//     <style>
//     </style>
//   </head>
//   <body>
//     <table>
//       <tbody>
//         <tr>
//           <td>
//             <a>
//               <img>
//               </img>
//             </a>
//             <br>
//             </br>
//             <div>
//               <a>
//                 <img>
//                 </img>
//               </a>
//               <a>
//                 <img>
//                 </img>
//               </a>
//               <a>
//                 <img>
//                 </img>
//               </a>
//             </div>
//             <br>
//             </br>
//           </td>
//           <td>
//             <h1>
//             </h1>
//             <p>
//               <br>
//               </br>
//               <br>
//               </br>
//               <br>
//               </br>
//               <tt>
//               </tt>
//               <tt>
//               </tt>
//               <tt>
//               </tt>
//             </p>
//             <div>
//               <table>
//                 <tbody>
//                   <tr>
//                     <td>
//                       <h1>
//                         <a>
//                         </a>
//                       </h1>
//                       <h1>
//                         <a>
//                         </a>
//                       </h1>
//                       <h1>
//                         <a>
//                         </a>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                     </td>
//                     <td>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                       </h1>
//                       <h1>
//                         <a>
//                         </a>
//                       </h1>
//                     </td>
//                   </tr>
//                   <tr>
//                     <td>
//                       <h1>
//                         <a>
//                         </a>
//                         <a>
//                         </a>
//                         <a>
//                         </a>
//                         <a>
//                         </a>
//                       </h1>
//                     </td>
//                   </tr>
//                 </tbody>
//               </table>
//             </div>
//             <p>
//               <a>
//                 <code>
//                 </code>
//               </a>
//               <a>
//                 <code>
//                 </code>
//               </a>
//               <a>
//                 <code>
//                 </code>
//               </a>
//               <a>
//                 <code>
//                 </code>
//               </a>
//             </p>
//             <p>
//               <a>
//               </a>
//               <a>
//               </a>
//             </p>
//           </td>
//         </tr>
//       </tbody>
//     </table>
//   </body>
// </html>
