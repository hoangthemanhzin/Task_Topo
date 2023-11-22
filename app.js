document.addEventListener("DOMContentLoaded", function () {
  //Tạo một thẻ SVG và đặt kích thước của nó.
  const width = 960,
      height = 500;

  const svg = d3.select("#network").append("svg")
      .attr("width", width)
      .attr("height", height);

  // Define your nodes and links (manually set positions for simplicity)
  let nodes = [
      // { id: 'Internet', group: 'internet', x: 100, y: 400 },
      // { id: 'e2', group: 'endpoint', x: 800, y: 400 },
      // { id: 'an1', group: 'an', x: 150, y: 400 },
      // { id: 'an2', group: 'an', x: 500, y: 400 },
      // { id: 'an3', group: 'an', x: 850, y: 400 },
      // { id: 'internet', group: 'internet'},
      // { id: 'e2', group: 'endpoint'},
      // { id: 'an1', group: 'an'},
      // { id: 'an2', group: 'an'},
      // { id: 'an3', group: 'an'},
      
      // UPF nodes will be placed later within the slices
      // { id: 'UPF1', group: 'upf' },
      // { id: 'UPF2', group: 'upf' },
      // { id: 'UPF3', group: 'upf' },
      // { id: 'UPF4', group: 'upf' },
      // { id: 'UPF5', group: 'upf' },
      
  ];

  let links = [
    // Links will be updated with positions after nodes are placed
    //{ source: 'internet', target: 'upf1', type: 'link' },
    //{ source: 'upf1', target: 'upf3', type: 'tran' },
    //{ source: 'upf3', target: 'upf2', type: 'tran' },
    //{ source: 'upf3', target: 'upf4', type: 'tran' },
    //{ source: 'e2', target: 'upf2', type: 'link' },
    //{ source: 'upf1', target: 'an1', type: 'link' },
    //{ source: 'upf2', target: 'an2', type: 'link' },
    //{ source: 'upf3', target: 'an2', type: 'link' },
    //{ source: 'upf4', target: 'an1', type: 'link' },
    //{ source: 'upf5', target: 'an2', type: 'link' },
    //{ source: 'upf5', target: 'upf1', type: 'tran' }
  ];

  // Định nghĩa các slice, mỗi slice chứa danh sách các UPF nodes và có các thuộc tính như vị trí và kích thước.
  const slices = [
      { id: 'slice1', upfs: ['UPF1', 'UPF3'], x: 50, y: 50, width: 300, height: 100 },
      { id: 'slice2', upfs: ['UPF2'], x: 400, y: 50, width: 150, height: 100 },
      { id: 'slice3', upfs: ['UPF4'], x: 600, y: 50, width: 150, height: 100 }
  ];

  //========================================================fetch api======================================================
  // Địa chỉ của API
  const apiUrl = 'http://127.0.0.1:9800/upmf2fe/topo';
  //let resultArrayTest = {}
  // Gọi API bằng fetch
  fetch(apiUrl)
    .then(response => {
      // Kiểm tra nếu có lỗi trong quá trình gọi API
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      // Chuyển đổi dữ liệu nhận được thành đối tượng JSON
      return response.json();
    })
    .then(data => {
      // Xử lý dữ liệu ở đây
      console.log("data for node : ", data.Nets);
      console.log("data for Infs : ", data.Nodes.upf2.Infs);

      // Thực hiện các thao tác khác với dữ liệu : 
      let resultArray = Object.entries(data.Nets).map(([id,group]) => ({
        id, 
        group : group === 0 ? 'an' : group === 1 ? 'tran' : 'internet' 
      }));
      console.log("data for result : ", resultArray)
      //resultArrayTest = resultArray
      nodes.push(...resultArray)
      console.log("nodes array : ", nodes)

      for(let nodeId in data.Nodes){
        if(data.Nodes.hasOwnProperty(nodeId)){
          const node = data.Nodes[nodeId]
          resultArray.push({
            id : node.Id,
            group : 'upf'
          })
        }
      }
      //Them upf vao nodes : 
      nodes.push(...resultArray)
      const linksArray = []
      for(let nodeId in data.Nodes){
        if(data.Nodes.hasOwnProperty(nodeId)){
          for(let netInfId in data.Nodes[nodeId].Infs){
            //console.log(data.Nodes[nodeId].Infs[test][0].Netname)
            if(data.Nodes[nodeId].Infs.hasOwnProperty(netInfId)){
              for(let index in data.Nodes[nodeId].Infs[netInfId]){
                console.log(data.Nodes[nodeId].Infs[netInfId][index])
                linksArray.push({
                  source : nodeId,
                  target : data.Nodes[nodeId].Infs[netInfId][index].Netname,
                  type   : 'link'
                })
              }
            }
          }
        }
      }
      links.push(...linksArray)
  //================================================================================================================
      // Sắp xếp danh sách links
      links.sort((a, b) => {
        if (a.target.group === 'upf' && b.target.group !== 'upf') {
          return 1;
        } else if (b.target.group === 'upf' && a.target.group !== 'upf') {
          return -1;
        } else {
          return 0;
        }
      });

      //Tạo các phần tử line để đại diện cho links.
      const linkElements = svg.selectAll(".link")
        .data(links)
        .enter().append("line")
        .attr("class", "link")
        .attr("fill", "none")
        .attr("stroke", "#ccc")

      //Tạo các phần tử circle để đại diện cho nodes.
      const nodeElements = svg.selectAll(".node")
        .data(nodes)
        .enter().append("circle")
        .attr("class", d => `node ${d.group}`)
        .attr("r", 20);

      //Tạo các phần tử text để hiển thị label cho các nodes.
      const textElements = svg.selectAll(".text-label")
        .data(nodes)
        .enter().append("text")
        .attr("class", "text-label")
        .text(d => d.id);

      //Thiết lập cơ chế kéo và thả cho nodes.
      const dragDrop = d3.drag()
        .on('start', node => {
            node.fx = node.x;
            node.fy = node.y;
        })
        .on('drag', node => {
            simulation.alphaTarget(0.7).restart();
            node.fx = d3.event.x;
            node.fy = d3.event.y;
        })
        .on('end', node => {
            if (!d3.event.active) {
                simulation.alphaTarget(0);
            }
            node.fx = null;
            node.fy = null;
        });

      nodeElements.call(dragDrop);
      textElements.call(dragDrop);

      // Create the simulation with forces (Tạo mô phỏng với các forces như link, charge, center, collision và slice.)
      const simulation = d3.forceSimulation(nodes)
        .force("link", d3.forceLink(links).id(d => d.id).distance(100))
        .force("charge", d3.forceManyBody().strength(-50))
        .force("center", d3.forceCenter(width/2, height /2))
        .force("collision", d3.forceCollide().radius(30))
        //.force("slice", forceSlice(slices))
        .on("tick", ticked);

      //Hàm này thực hiện cập nhật vị trí của các UPF nodes để chúng nằm trong các slice.
      function forceSlice(slices) {
        const force = function (alpha) {
          nodes.forEach(function (d) {
              if (d.group === 'upf') {
                  const slice = slices.find(s => s.upfs.includes(d.id));
                  const padding = 20; // Padding from the edge of the slice
                  d.x = Math.max(slice.x + padding, Math.min(slice.x + slice.width - padding, d.x));
                  d.y = Math.max(slice.y + padding, Math.min(slice.y + slice.height - padding, d.y));
              }
          });
      };
      return force;
      }

      //Hàm này được gọi mỗi lần simulation được cập nhật, cập nhật vị trí của các phần tử link, node, và text.
      function ticked() {
        linkElements
            .attr("x1", d => d.source.x)
            .attr("y1", d => d.source.y)
            .attr("x2", d => d.target.x)
            .attr("y2", d => d.target.y);

        nodeElements
            .attr("cx", d => d.x)
            .attr("cy", d => d.y);

        textElements
            .attr("x", d => d.x)
            .attr("y", d => d.y + 35);
      }

      // Hàm tạo lực của link được cập nhật
      const linkForce = d3.forceLink(links)
        .id(d => d.id)
        .distance(50)
        .iterations(1)
        .links(links);

      // Thiết lập lực của link trong simulation
      simulation.force("link", linkForce);

    })
    .catch(error => {
      // Xử lý lỗi nếu có
      console.error('Error fetching data:', error.message);
    });

  //=======================================================================================================================

  // Vẽ các hình chữ nhật đại diện cho các slice.
  // slices.forEach(slice => {
  //   svg.append("rect")
  //       .attr("x", slice.x)
  //       .attr("y", slice.y)
  //       .attr("width", slice.width)
  //       .attr("height", slice.height)
  //       .attr("class", "upf-rect");
  // });


//   // Sắp xếp danh sách links
//   links.sort((a, b) => {
//     if (a.target.group === 'upf' && b.target.group !== 'upf') {
//       return 1;
//     } else if (b.target.group === 'upf' && a.target.group !== 'upf') {
//       return -1;
//     } else {
//       return 0;
//     }
//   });

//   //Tạo các phần tử line để đại diện cho links.
//   const linkElements = svg.selectAll(".link")
//     .data(links)
//     .enter().append("line")
//     .attr("class", "link")
//     .attr("fill", "none")
//     .attr("stroke", "#ccc")

//   //Tạo các phần tử circle để đại diện cho nodes.
//   const nodeElements = svg.selectAll(".node")
//     .data(nodes)
//     .enter().append("circle")
//     .attr("class", d => `node ${d.group}`)
//     .attr("r", 20);

//   //Tạo các phần tử text để hiển thị label cho các nodes.
//   const textElements = svg.selectAll(".text-label")
//     .data(nodes)
//     .enter().append("text")
//     .attr("class", "text-label")
//     .text(d => d.id);

//   //Thiết lập cơ chế kéo và thả cho nodes.
//   const dragDrop = d3.drag()
//     .on('start', node => {
//         node.fx = node.x;
//         node.fy = node.y;
//     })
//     .on('drag', node => {
//         simulation.alphaTarget(0.7).restart();
//         node.fx = d3.event.x;
//         node.fy = d3.event.y;
//     })
//     .on('end', node => {
//         if (!d3.event.active) {
//             simulation.alphaTarget(0);
//         }
//         node.fx = null;
//         node.fy = null;
//     });

//   nodeElements.call(dragDrop);
//   textElements.call(dragDrop);

//   // Create the simulation with forces (Tạo mô phỏng với các forces như link, charge, center, collision và slice.)
//   const simulation = d3.forceSimulation(nodes)
//     .force("link", d3.forceLink(links).id(d => d.id).distance(100))
//     .force("charge", d3.forceManyBody().strength(-50))
//     .force("center", d3.forceCenter(width/2, height /2))
//     .force("collision", d3.forceCollide().radius(30))
//     //.force("slice", forceSlice(slices))
//     .on("tick", ticked);

//   //Hàm này thực hiện cập nhật vị trí của các UPF nodes để chúng nằm trong các slice.
//   function forceSlice(slices) {
//     const force = function (alpha) {
//       nodes.forEach(function (d) {
//           if (d.group === 'upf') {
//               const slice = slices.find(s => s.upfs.includes(d.id));
//               const padding = 20; // Padding from the edge of the slice
//               d.x = Math.max(slice.x + padding, Math.min(slice.x + slice.width - padding, d.x));
//               d.y = Math.max(slice.y + padding, Math.min(slice.y + slice.height - padding, d.y));
//           }
//       });
//   };
//   return force;
//   }

//   //Hàm này được gọi mỗi lần simulation được cập nhật, cập nhật vị trí của các phần tử link, node, và text.
//   function ticked() {
//     linkElements
//         .attr("x1", d => d.source.x)
//         .attr("y1", d => d.source.y)
//         .attr("x2", d => d.target.x)
//         .attr("y2", d => d.target.y);

//     nodeElements
//         .attr("cx", d => d.x)
//         .attr("cy", d => d.y);

//     textElements
//         .attr("x", d => d.x)
//         .attr("y", d => d.y + 35);
//   }

//   // Hàm tạo lực của link được cập nhật
//   const linkForce = d3.forceLink(links)
//     .id(d => d.id)
//     .distance(50)
//     .iterations(1)
//     .links(links);

//   // Thiết lập lực của link trong simulation
//   simulation.force("link", linkForce);

//   // Add labels to slices
//   // slices.forEach(slice => {
//   //   svg.append("text")
//   //       .attr("x", slice.x + slice.width / 2)
//   //       .attr("y", slice.y + slice.height + 20)
//   //       .attr("text-anchor", "middle")
//   //       .text(slice.id);
//   // });
});

// // fetch api : 
