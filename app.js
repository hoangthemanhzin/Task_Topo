document.addEventListener("DOMContentLoaded", function () {
  const width = 960,
      height = 500;

  const svg = d3.select("#network").append("svg")
      .attr("width", width)
      .attr("height", height);

  let nodes = [];
  let links = [];

  const slices = [
      { id: 'slice1', upfs: ['UPF1', 'UPF3'], x: 50, y: 50, width: 300, height: 100 },
      { id: 'slice2', upfs: ['UPF2'], x: 400, y: 50, width: 150, height: 100 },
      { id: 'slice3', upfs: ['UPF4'], x: 600, y: 50, width: 150, height: 100 }
  ];

  const apiUrl = 'http://127.0.0.1:9800/upmf2fe/topo';

  fetch(apiUrl)
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
      const additionalButtonsContainer = d3.select("#additionalButtons");
      let savedButtons = "";
      for(let i = 0; i < localStorage.length; i++){
        const key = localStorage.key(i);
        const value = localStorage.getItem(key);
        savedButtons = savedButtons + value
        console.log("check : ", savedButtons)
      }
      if (savedButtons) {
        additionalButtonsContainer.html(savedButtons);
        //savedButtons = "";
      }
    nodes.push(...processNetworkData(data.Nets));
    nodes.push(...processNodeData(data.Nodes));
    links.push(...processLinkData(data.Nodes));
    links.sort((a, b) => {
      if (a.target.group === 'upf' && b.target.group !== 'upf') {
        return 1;
      } else if (b.target.group === 'upf' && a.target.group !== 'upf') {
        return -1;
      } else {
        return 0;
      }
    });

    const simulation = d3.forceSimulation(nodes)
      .force("link", d3.forceLink(links).id(d => d.id).distance(100))
      .force("charge", d3.forceManyBody().strength(-50))
      .force("center", d3.forceCenter(width/2, height /2))
      .force("collision", d3.forceCollide().radius(30))
      .on("tick", ticked);

    console.log("node : ", nodes)
    console.log("links : ", links)
    drawElements(svg, links, nodes);
    setDragDrop(simulation);
    setLinkForce(simulation, links);

  })
  .catch(error => {
    console.error('Error fetching data:', error.message);
  });



  function processNetworkData(networkData) {
    return Object.entries(networkData).map(([id, group]) => ({
      id, 
      group: group === 0 ? 'an' : group === 1 ? 'tran' : 'internet' 
    }));
  }

  function processNodeData(nodeData) {
    const resultArray = [];
    for (let nodeId in nodeData) {
      if (nodeData.hasOwnProperty(nodeId)) {
        const node = nodeData[nodeId];
        resultArray.push({ id: node.Id, group: 'upf' });
      }
    }
    return resultArray;
  }

  function processLinkData(nodeData) {
    const linksArray = [];
    for (let nodeId in nodeData) {
      if (nodeData.hasOwnProperty(nodeId)) {
        for (let netInfId in nodeData[nodeId].Infs) {
          if (nodeData[nodeId].Infs.hasOwnProperty(netInfId)) {
            for (let index in nodeData[nodeId].Infs[netInfId]) {
              linksArray.push({
                source: nodeId,
                target: nodeData[nodeId].Infs[netInfId][index].Netname,
                type: 'link'
              });
            }
          }
        }
      }
    }
    return linksArray;
  }

  function drawElements(svg, links, nodes) {
    const linkedNodes = Array.from(new Set(links.flatMap(link => [link.source.id, link.target.id])));
    svg.selectAll(".link")
      .data(links)
      .enter().append("line")
      .attr("class", "link")
      .attr("fill", "none")
      .attr("stroke", "#ccc");

    const nodeElements = svg.selectAll(".node")
      .data(nodes.filter(node => linkedNodes.includes(node.id)))
      .enter().append("circle")
      .attr("class", d => `node ${d.group}`)
      .attr("r", 20);

    svg.selectAll(".text-label")
      .data(nodes.filter(node => linkedNodes.includes(node.id)))
      .enter().append("text")
      .attr("class", "text-label")
      .text(d => d.id);

    // Trong phần xử lý sự kiện click trên node
    nodeElements.on('click', handleNodeClick);
    
  }

  function setDragDrop(simulation) {
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

    svg.selectAll(".node").call(dragDrop);
    svg.selectAll(".text-label").call(dragDrop);
  }

  function setLinkForce(simulation, links) {
    const linkForce = d3.forceLink(links)
      .id(d => d.id)
      .distance(50)
      .iterations(1)
      .links(links);

    simulation.force("link", linkForce);
  }

  function ticked() {
    svg.selectAll(".link")
      .attr("x1", d => d.source.x)
      .attr("y1", d => d.source.y)
      .attr("x2", d => d.target.x)
      .attr("y2", d => d.target.y);

    svg.selectAll(".node")
      .attr("cx", d => d.x)
      .attr("cy", d => d.y);

    svg.selectAll(".text-label")
      .attr("x", d => d.x)
      .attr("y", d => d.y + 35);
  }

  // Khai báo biến selectedNode để lưu trữ node đã được chọn
  let selectedNode;

  // Hàm xử lý sự kiện click trên node
  function handleNodeClick(d) {
    if(d.group == 'upf'){
    // Ẩn tất cả các button trước khi hiển thị mới
    hideAllButtons();

    // Hiển thị hai button
    d3.select("#deactivateBtn").style("display", "block");

    // Lưu trữ thông tin về node đã được chọn
    selectedNode = d;
    console.log("node selected :", selectedNode)
    }

  }

  // Hàm ẩn tất cả các button
  function hideAllButtons() {
    d3.select("#deactivateBtn").style("display", "none");
  }

  function handleActionButtonClick(action) {
    if (selectedNode) {
      console.log(`${action} node: `, selectedNode.id);
      const additionalButtonsContainer = d3.select("#additionalButtons");
      additionalButtonsContainer.append("button")
      .attr("class", "additionalBtn")
      .attr("id", `${selectedNode.id}`)
      .text(`Activate node ${selectedNode.id}`)
      .style("display", "block");
      sendPostDeactivateRequest(selectedNode.id)
      // Xem them ve localStorage : 
      localStorage.setItem(`${selectedNode.id}`, ` <button id="${selectedNode.id}" style="display: block;">Activate node ${selectedNode.id}</button>`);
      hideAllButtons()
      // Logic chuyen mau node tu trang sang den : 
      // Ham logic sua lai cac link thanh net dut : 
    }
  }

  // Trong phần xử lý sự kiện click cho button "Deactivate"
  d3.select("#deactivateBtn").on('click', () => handleActionButtonClick("Deactivating"));

  d3.select("#additionalButtons").on('click', function () {
    // Lấy thông tin về nút được click
    const clickedButton = d3.event.target;
    const buttonId = clickedButton.id;

    // Gọi hàm xử lý sự kiện cho nút bổ sung
    handleAdditionalButtonClick(buttonId);
  });

  function handleAdditionalButtonClick(buttonId) {
    // Thực hiện hành động tương ứng với nút được click
    console.log(`${buttonId}`);
    // Thêm logic xử lý tương ứng ở đây : 
    sendPostActivateRequest(buttonId)
    d3.select("#activateBtn").style("display", "none");
    localStorage.removeItem(`${buttonId}`)
  }
  
  function sendPostDeactivateRequest(upfId) {
    const apiUrl = 'http://127.0.0.1:9800/upmf2activate/deactivate';
    const requestData = {
      UpfIds: [upfId]
    };
  
    fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestData)
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json();
    })
    .then(data => {
      console.log('Deactivation successful:', data);
    })
    .catch(error => {
      console.error('Error deactivating:', error.message);
    });
  }

  function sendPostActivateRequest(upfId) {
    const apiUrl = 'http://127.0.0.1:9800/upmf2activate/activate';
    const requestData = {
      UpfIds: [upfId]
    };
  
    fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestData)
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json();
    })
    .then(data => {
      console.log('Deactivation successful:', data);
    })
    .catch(error => {
      console.error('Error deactivating:', error.message);
    });
  }
});
